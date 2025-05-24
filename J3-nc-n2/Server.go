package main

import (
	"net"
	"fmt"
)

func NewServer(addr string,exec string) {
	listener, err := net.Listen("tcp", addr)
	CheckError(err,"NewServer-1")
	/*代替了👆
	if err!=nil{
		fmt.Println(err)
	}
	*/
	defer listener.Close()

	Conn, err := listener.Accept()
	CheckError(err,"NewServer-2")
	defer Conn.Close()


	if exec!=""{
		ExecBindToConn(Conn,exec)
	}else{
		go ReadFromConn(Conn)
		WriteToConn(Conn)
	}



/*	done := make(chan struct{})
	CopyToConn(Conn,done)
	GetFromConn(Conn,done)

	<-done // 任意一方出错，关闭连接*/
	fmt.Println("[*] Connection closed")

}
