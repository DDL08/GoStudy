package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {

	//cmd := exec.Command("ipconfig") // Windows
	cmdStr := flag.String("cmd", "", "无")
	flag.Parse()

	if *cmdStr == "" {
		fmt.Println("error1")
		return
	}

	cmd := exec.Command("cmd", "/c", *cmdStr)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("命令执行出错:", err)
	}

	// 3. 打印命令输出
	fmt.Println("命令输出：")
	fmt.Println(string(output))
}
