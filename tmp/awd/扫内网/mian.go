package main

import(
	"fmt"
	"net/http"
	"sync"
	"os"
)
func main(){

	file,err:=os.Create("C:/Users/LEGION/Desktop/host1.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	var mu sync.Mutex //写文件时加锁

	client:=&http.Client{}

	for i:=1;i<254;i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			url:=fmt.Sprintf("http://192-168-1-%d.pvp6006.bugku.cn", i)
			resp,err:=
		}
	}


}