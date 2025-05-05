// 实现一个web服务
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", HelloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "url:=%q\n", req.URL.Path)
}
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}

}
