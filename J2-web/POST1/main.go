package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://github.com/DDL08/GoStudy/blob/main/1webGee/base1/main1.go"
	payload := "username=1&password=2"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(payload))
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("响应代码%d\n", resp.StatusCode)
	//fmt.Println("Response Body: \n%s\n", body)
	fmt.Printf("Response Body: \n%s\n", body)
	//Println不支持占位符
}
