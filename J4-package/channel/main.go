package main

import (
	"fmt"
	"time"
)

func worker(c chan bool) {
	fmt.Println("任务处理中...")
	time.Sleep(2 * time.Second)
	c <- true // 通知主线程：完成了
}

func main() {
	done := make(chan bool)
	go worker(done)

	<-done // 等待 worker 通知
	fmt.Println("主线程收到通知：任务已完成")
}
