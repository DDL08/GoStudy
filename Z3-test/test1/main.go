package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	// 假设 tokens[0] 是你想执行的命令
	//exec.Command("cmd.exe", "/c", "chcp 65001").Run()
	tokens := []string{"cmd.exe", "/c", "ipconfig"} // 这里可以替换为你想测试的命令

	// 创建 exec.Command 来执行命令

	var cmd *exec.Cmd
	if len(tokens) == 1 {
		cmd = exec.Command(tokens[0])
	} else {
		cmd = exec.Command(tokens[0], tokens[1:]...)
	}

	// 捕获输出
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 将输出从 GBK 转换为 UTF-8

	decoder := simplifiedchinese.GBK.NewDecoder()
	utf8Output, err := ioutil.ReadAll(transform.NewReader(&out, decoder))
	// 执行命令

	// 打印命令输出
	fmt.Println("Command Output:")
	fmt.Println(string(utf8Output))
}
