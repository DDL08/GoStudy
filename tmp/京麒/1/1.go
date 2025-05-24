package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// 连接 Node.js 服务端（假设在 localhost:8081）
	conn, err := net.Dial("tcp", "39.106.16.204:59101")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()

	// 要访问的 URL（应为 http 或 https 开头）
	url := "http://localhost:11016" // 请换成你实际要测试的 URL
	fmt.Fprintf(conn, "%s\n", url)

	// 读取服务端返回内容
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("服务器返回:", scanner.Text())
	}
}
