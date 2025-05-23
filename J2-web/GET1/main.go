package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://github.com/DDL08/GoStudy/blob/main/1webGee/base1/main1.go"
	resp, err := http.Get(url)
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
