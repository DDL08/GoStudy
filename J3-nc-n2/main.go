package main

import (
	"flag"
	"fmt"
)

func main() {
	listen := flag.Bool("l", false, "Listen Mode")
	//调用 flag 包的方法，定义一个 bool 类型的命令行参数。(参数，默认值，help语句)
	port := flag.String("p", "444", "port you used")
	host := flag.String("h", "127.0.0.1", "will be connected")
	exec := flag.String("e", "", "Program to execute after connection") // 添加 -e 参数
	flag.Parse()

	address := *host + ":" + *port

	if *listen {
		fmt.Println("[*] Listening on", address)
		NewServer(address, *exec)

	} else {
		fmt.Println("[*] Connecting to", address)
		NewClient(address, *exec)
	}

}
