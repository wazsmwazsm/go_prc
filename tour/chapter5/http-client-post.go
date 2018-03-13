package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	// 构造请求参数
	v := url.Values{}
	v.Set("username", "admin")
	v.Set("password", "123456")

	// 构造请求主体
	request_body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	// 构建客户端对象
	client := &http.Client{}
	path := "http://community_back_end.app/api/auth/login"
	request, err := http.NewRequest("POST", path, request_body)
	if err != nil {
        fmt.Println("Fatal error ", err.Error())
	}
	// 设置头信息
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 执行请求
	response, err := client.Do(request)
	if err != nil {
        fmt.Println("Fatal error ", err.Error())
	}
	
	defer response.Body.Close()

	// 打印结果
	response_body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(response_body))
}
