package main

import (
	"fmt"
	"os"
)

func main() {
	dic := "abcdefg"
	outputfile := "F:\\addl\\go\\GoStudy\\J1-io\\Write_world4-wordlist\\dict.txt"

	file, err := os.Create(outputfile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for a := 0; a < len(dic); a++ {
		for b := 0; b < len(dic); b++ {
			for c := 0; c < len(dic); c++ {
				for d := 0; d < len(dic); d++ {
					for e := 0; e < len(dic); e++ {
						strings := string(dic[a]) + string(dic[b]) + string(dic[c]) + string(dic[d]) + string(dic[e])
						file.WriteString(strings + "\n")

					}
				}
			}
		}
	}

	fmt.Println("任务完成")
}
