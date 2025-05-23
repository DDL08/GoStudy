package main

import (
	"fmt"
	"os"
)

func main() {
	a, err := os.ReadFile("F:\\addl\\go\\GoStudy\\J1-io\\rw1\\1.txt")
	if err == nil {
		fmt.Println(err)
	}
	err = os.WriteFile("F:\\addl\\go\\GoStudy\\J1-io\\rw1\\2.txt", a, 0644)
}
