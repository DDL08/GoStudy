package main

import (
	"fmt"
	"os"
)

func main() {

	//dic:="ZXCVBNMLKJHGFDSAQWERTYUIOPabcdefghijklmnopqrstuvwxyz1234567890"
	//dic2:="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,."
	dic3 := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	outputfile := "F:\\addl\\go\\GoStudy\\J1-io\\Write_world3-wordlist\\dict.txt"

	file, err := os.Create(outputfile)
	if err != nil {
		fmt.Println("gogo", err)
	}
	defer file.Close()
	for _, a := range dic3 {
		for _, b := range dic3 {
			for _, c := range dic3 {
				for _, d := range dic3 {
					for _, e := range dic3 {
						strings := string(a) + string(b) + string(c) + string(d) + string(e)
						file.WriteString(strings + "\n")
					}
				}
			}
		}

	}
	fmt.Println("保存完毕")

}
