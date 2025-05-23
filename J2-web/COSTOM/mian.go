package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://github.com/DDL08/GoStudy/blob/main/1webGee/base1/main1.go"
	//payload := "username=1"
	//req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	req, err := http.NewRequest("GET", url, nil) //, strings.NewReader(payload))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Go-http-client/1.1")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("响应头:%d", resp.StatusCode)
	fmt.Printf("文本：\n%s\n", body)
}
