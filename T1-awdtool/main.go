package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var flagRegex = regexp.MustCompile(`flag\{[a-zA-Z0-9_-]+\}`)

func sendRequest(target, body, method string, headers map[string]string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(method, target, bytes.NewBufferString(body))
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data), nil
}

func extractFlag(resp string) string {
	match := flagRegex.FindString(resp)
	return match
}

func submitFlag(api, flag, token string) (string, error) {
	data := fmt.Sprintf("token=%s&flag=%s", token, flag)
	resp, err := http.Post(api, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("AWD Flag Submit GUI")

	ipList := []string{}

	apiEntry := widget.NewEntry()
	tokenEntry := widget.NewEntry()
	methodSelect := widget.NewSelect([]string{"GET", "POST"}, nil)
	methodSelect.SetSelected("GET")
	customHeader := widget.NewMultiLineEntry()
	postBody := widget.NewMultiLineEntry()
	concurrencyEntry := widget.NewEntry()
	concurrencyEntry.SetText("5")

	logOutput := widget.NewMultiLineEntry()
	logOutput.SetMinRowsVisible(10)

	loadFileBtn := widget.NewButton("导入 IP 列表 (CSV/TXT)", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if err != nil || uri == nil {
				return
			}
			defer uri.Close()
			scanner := bufio.NewScanner(uri)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line != "" {
					ipList = append(ipList, line)
				}
			}
			logOutput.SetText(fmt.Sprintf("已导入 %d 个地址\n", len(ipList)))
		}, myWindow)
	})

	runBtn := widget.NewButton("开始运行", func() {
		concurrency, _ := strconv.Atoi(concurrencyEntry.Text)
		sem := make(chan struct{}, concurrency)
		var wg sync.WaitGroup
		logFile, _ := os.Create("log.txt")
		defer logFile.Close()

		headers := map[string]string{}
		headerLines := strings.Split(customHeader.Text, "\n")
		for _, h := range headerLines {
			parts := strings.SplitN(h, ":", 2)
			if len(parts) == 2 {
				headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}

		for _, target := range ipList {
			wg.Add(1)
			sem <- struct{}{}
			go func(t string) {
				defer wg.Done()
				defer func() { <-sem }()
				logLine := fmt.Sprintf("访问: %s\n", t)
				resp, err := sendRequest(t, postBody.Text, methodSelect.Selected, headers)
				if err != nil {
					logLine += fmt.Sprintf("请求失败: %v\n", err)
				} else {
					flag := extractFlag(resp)
					if flag == "" {
						logLine += "未找到 flag\n"
					} else {
						logLine += fmt.Sprintf("提取到 flag: %s\n", flag)
						result, err := submitFlag(apiEntry.Text, flag, tokenEntry.Text)
						if err != nil {
							logLine += fmt.Sprintf("提交失败: %v\n", err)
						} else {
							logLine += fmt.Sprintf("提交成功: %s\n", result)
						}
					}
				}
				logOutput.SetText(logOutput.Text + logLine)
				logFile.WriteString(logLine)
			}(target)
		}
		wg.Wait()
		logOutput.SetText(logOutput.Text + "\n全部完成\n")
	})

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("提交接口地址:"), apiEntry,
		widget.NewLabel("认证 Token:"), tokenEntry,
		widget.NewLabel("请求方法:"), methodSelect,
		widget.NewLabel("自定义 Header (每行一个 key:value):"), customHeader,
		widget.NewLabel("POST 请求 Body (GET 可忽略):"), postBody,
		widget.NewLabel("并发数限制:"), concurrencyEntry,
		loadFileBtn,
		runBtn,
		widget.NewLabel("运行日志输出:"), logOutput,
	))

	myWindow.Resize(fyne.NewSize(800, 700))
	myWindow.ShowAndRun()
}
