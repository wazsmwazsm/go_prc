package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func main() {
	// 设置路由和处理回调
	http.HandleFunc("/", sayHelloName)
	// 启动服务, 监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析请求参数，默认不会解析
	fmt.Println(r.Form) 
	fmt.Println("path", r.URL.Path) 
	fmt.Println("scheme", r.URL.Scheme) 
	fmt.Println(r.Form["url_long"]) 
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, " "))
	}
	// 写入响应
	fmt.Fprintf(w, "Hello astaxie!")
}
