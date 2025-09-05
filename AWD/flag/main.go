package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

const RequestMethod = "GET"

// POST请求体，GET请求时可留空
const PostBody = "param1=value1&param2=value2"

// 提交flag的接口地址
const SubmitAPI = "https://example.com/submit"

// 认证token
const Token = "your_token_here"

// 自定义请求头（key:value），不需要可以留空
var CustomHeaders = map[string]string{
	"User-Agent": "GoClient",
	// "Authorization": "Bearer token",
}
/*------------------------------发送请求正则flag('GET+POST')↓*/
var flagRegex = regexp.MustCompile(`flag\{[a-zA-Z0-9_-]+\}`)

func sendRequest(url string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	var reqBody *bytes.Buffer
	if strings.ToUpper(RequestMethod) == "POST" {
		reqBody = bytes.NewBufferString(PostBody)
	} else {
		reqBody = nil
	}

	req, err := http.NewRequest(RequestMethod, url, reqBody)
	CheckError(err,"")

	for k, v := range CustomHeaders {
		req.Header.Set(k, v)
	}

	// POST请求常用content-type
	if strings.ToUpper(RequestMethod) == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(req)
	CheckError(err,"")
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	return string(data), nil
}

func extractFlag(resp string) string {
	return flagRegex.FindString(resp)
}
/*------------------------------提交flag模块↓*/
func submitGetFlag(flag string) (string, error) {
	url := fmt.Sprintf("https://example.com/submit?token=%s&flag=%s", Token, flag)
	resp, err := http.Get(url)
	CheckError(err, "submitGetFlag-Error")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
func submitPostFlag(flag string) (string, error) {
	data := fmt.Sprintf("token=%s&flag=%s", Token, flag)
	resp, err := http.Post(SubmitAPI, "application/x-www-form-urlencoded", strings.NewReader(data))
	CheckError(err, "submitPostFlag-Error")
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}
/*------------------------------checkerror↓*/
func CheckError(err error, str string) {
	if err != nil {
		fmt.Println(str)
		return
	}

}
/*------------------------------*/
func main() {
	// 从文件读URL列表，每行一个
	file, err := os.Open("C:/Users/LEGION/Desktop/host1.txt")
	CheckError(err, "打开iplist.txt失败:")
	defer file.Close()

	//把IP塞入数组
	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, 5) // 并发数5，可修改

	for _, url := range urls {
		wg.Add(1)
		sem <- struct{}{}

		go func(u string) {
			defer wg.Done()
			defer func() { <-sem }()

			fmt.Printf("访问: %s\n", u)
			resp, err := sendRequest(u)
			CheckError(err, "请求失败:")

			flagStr := extractFlag(resp)
			if flagStr == "" {
				fmt.Println("未找到flag")
				return
			}

			fmt.Printf("提取flag: %s\n", flagStr)
			res, err := submitGetFlag(flagStr)
			CheckError(err, "提交失败")

			fmt.Printf("提交结果: %s\n", res)
		}(url)
	}

	wg.Wait()
	fmt.Println("全部完成")
}
