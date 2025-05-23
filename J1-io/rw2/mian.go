package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("gogogo", err)
	}
	input = strings.TrimSpace(input)
	////这个输出回车也会被算(不添加这个的话)
	fmt.Println(input)

}
