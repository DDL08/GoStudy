package main

import (
	"fmt"
	"os"
)

func main() {

	usernames := []string{"admin", "DDL", "user", "test"}
	suffixes := []string{"123", "456", "789"}
	outputfile := "F:\\addl\\go\\GoStudy\\J1-io\\Write_world1\\dict.txt"
	//outputfile := "F:/addl/go/GoStudy/J1-io/Write_world1/main.go/dict.txt"
	//这Linux路径，Windows报错，Windows的双反斜线转义
	file, err := os.Create(outputfile)

	if err != nil {
		fmt.Println("", err)
	}

	defer file.Close()

	for _, user := range usernames {
		for _, suf := range suffixes {
			word := user + suf
			file.WriteString(word + "\n")
		}
	}
}
