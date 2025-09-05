package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"Z2/protobufapi"

	"log"
	"os"
	"Z2/util"
)
//	"gotoexec/util"
func main()  {

	var (
		opts []grpc.DialOption
		conn *grpc.ClientConn
		err error
		client protobufapi.ClientClient
		session, ip  string
		sleepTime, port int

	)
	flag.IntVar(&sleepTime,"sleep",0,"sleep time")
	flag.StringVar(&session,"session","","start session")
	flag.StringVar(&ip,"ip","127.0.0.1","Server IP")
	flag.IntVar(&port,"port",1962,"AdminServer Port")
	flag.Parse()


	// WithInsecure 忽略证书
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 12 )))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(1024 * 1024 * 12)))
	if conn,err = grpc.Dial(fmt.Sprintf("%s:%d",ip, port),opts...);
	err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client = protobufapi.NewClientClient(conn)

	if sleepTime != 0 {
		var time = new(protobufapi.SleepTime)
		time.Time = int32(sleepTime)
		ctx := context.Background()
		client.SetSleepTime(ctx,time)
	}

	if session != "" {
		if session == "start" {
			fmt.Println("start exec:")
			for {
				var cmd = new(protobufapi.Command)
				//go中scan、scanf、scanln在输入时都会将空格作为一个字符串的结束
				//fmt.Scan(&command)
				reader := bufio.NewReader(os.Stdin)
				command, _, err := reader.ReadLine()
				if nil != err {
					fmt.Println("reader.ReadLine() error:", err)
				}
				cmd = Run(cmd,command,client)
				out:= cmd.Out
				if err != nil {
					log.Fatal(err.Error())
				}
				cmd.Out = util.ConvertByte2String([]byte (out), util.GB18030)
				fmt.Println(cmd.Out)
			}
		} else {
			fmt.Println("please input start")
		}
	}


}

func Run(cmd *protobufapi.Command,command []byte,client protobufapi.ClientClient) *protobufapi.Command{
	var err error
	cmd.In = string(command)
	ctx := context.Background()
	//x := *client
	cmd, err =client.RunCommand(ctx, cmd)
	if err != nil {
		log.Fatal("client"+err.Error())
	}
	return cmd
	//result,_ := util.DecryptByAes(cmd.Out)
}