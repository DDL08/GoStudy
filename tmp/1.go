package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, world!", i)
		time.Sleep(1 * time.Second)
	}
}
