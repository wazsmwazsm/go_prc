package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// TIMEOUT 超时时间
const TIMEOUT = 4

func main() {
	for {
		chData := make(chan []byte)
		timeout := time.NewTimer(time.Duration(TIMEOUT) * time.Second)

		go request(chData) // run request

		select {
		case <-timeout.C:
			fmt.Println("Current request timeout, begin next request")
		case data := <-chData:
			fmt.Println(string(data))
		}

	}
}

// request 请求消息
func request(ch chan []byte) {

	resp, err := http.Get("http://127.0.0.1:6006")
	if err != nil {
		fmt.Println("Request faild!")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 504 { // 服务超时
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body read faild!")
		return
	}

	ch <- body
}
