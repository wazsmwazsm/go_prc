package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func serverInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数 (默认不解析)
	// 打印服务端信息
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	// 打印参数
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // 输出到客户端的信息
}

func main() {
	http.HandleFunc("/", serverInfo)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
