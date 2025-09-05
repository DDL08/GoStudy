package main

import (
	"os/signal"
	"runtime"
	"syscall"

	"github.com/DDL08/GoStudy/v2/cli"
	"github.com/DDL08/GoStudy/v2/crypto"
	"github.com/DDL08/GoStudy/v2/dispather"
	"github.com/DDL08/GoStudy/v2/node"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cli.ParseArgs()
	// ignore sighup
	signal.Ignore(syscall.SIGHUP)
	//在终端关闭后依然运行

	node.CurrentNode.IsAdmin = 0
	crypto.InitEncryption(cli.Args.Password)

	node.CurrentNode.InitCommandBuffer()
	node.CurrentNode.InitDataBuffer()

	dispather.InitAgentHandler()

	initnode.InitNode()
}
