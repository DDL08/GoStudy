package main

import (
	"fmt"
	"net"
)

func NewClient(addr string, exec string) {
	Conn, err := net.Dial("tcp", addr)
	CheckError(err, "NewClient")
	defer Conn.Close()

	if exec != "" {
		ExecBindToConn(Conn, exec)
	} else {
		go ReadFromConn(Conn)
		WriteToConn(Conn)
	}
	/*
		done := make(chan struct{})
		CopyToConn(Conn, done)
		GetFromConn(Conn, done)

		<-done // 任意一方出错，关闭连接*/
	fmt.Println("[*] Connection closed")

}
