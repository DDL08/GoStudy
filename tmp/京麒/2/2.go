package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://39.106.16.204:11016/", nil)
	if err != nil {
		panic(err)
	}

	// 手动设置cookie
	req.Header.Set("Cookie", "flag=flag{your_flag_here}")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("响应状态:", resp.Status)
	fmt.Println("响应内容:")
	fmt.Println(string(body))
}
