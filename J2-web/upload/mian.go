package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const fileName = "data.txt"

func uploadFile() {
	for i := 0; i < 5; i++ {
		content := fmt.Sprintf("upload %d at %s\n", i, time.Now().Format(time.RFC3339))
		err := ioutil.WriteFile(fileName, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
		} else {
			fmt.Println("[UPLOAD] Wrote:", content)
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func useFile() {
	for i := 0; i < 20; i++ {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("[USE] Error reading file:", err)
		} else {
			fmt.Printf("[USE] Read (%d): %s", i, string(data))
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 创建文件以避免首次读取错误
	_ = ioutil.WriteFile(fileName, []byte("initial content\n"), 0644)

	go uploadFile()
	go useFile()

	// 等待所有 goroutine 完成
	time.Sleep(3 * time.Second)

	// 清理文件
	_ = os.Remove(fileName)
}
