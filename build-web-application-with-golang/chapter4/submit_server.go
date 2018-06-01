package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	http.HandleFunc("/", sayhelloname)
	http.HandleFunc("/submit", submit)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func sayhelloname(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // request.Form 是一个 url.Values 类型, 可以用 Set、Add 方法设置、添加内容
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world!") 
}

func submit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("submit.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		
		// 为空验证

		// if len(r.Form["username"][0]) == 0 { // 这种可以取 map 值
		// r.Form.Get()只能获取单个的值
		if len(r.Form.get("username")) == 0 { 
			fmt.Fprintf(w, "用户名为空！") 
		}

		if len(r.Form.get("password")) == 0 {
			fmt.Fprintf(w, "密码为空！") 
		}

		// 数字验证
		age, err := strconv.Atoi(r.Form.get("age")) 
		if err != nil {
			fmt.Fprintf(w, "年龄必须是数字") 
		}

		if age > 100 || age <= 0 {
			fmt.Fprintf(w, "年龄范围必须在 1~100") 
		}

		// 

	}
}
