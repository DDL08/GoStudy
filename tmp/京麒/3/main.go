package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://39.106.16.204:11016/" // 替换成你要请求的URL

	// 新建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// 可以设置cookie，比如发送flag cookie
	req.Header.Set("Cookie", "flag=flag{your_flag_here}")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("响应状态:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("响应内容:")
	fmt.Println(string(body))
}
