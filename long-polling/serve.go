package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// TIMEOUT 服务端超时时间
const TIMEOUT = 20

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("Serve start...")
	http.ListenAndServe(":6006", nil)
}

// Handler 处理请求
func Handler(w http.ResponseWriter, r *http.Request) {

	chData := make(chan []byte)
	timeout := time.NewTimer(time.Duration(TIMEOUT) * time.Second)

	go Service(chData) // run service

	// select 阻塞 goroutine, hold TCP 连接
	select {
	case <-timeout.C:
		w.WriteHeader(503)
		fmt.Fprintln(w, "Timeout")
	case data := <-chData:
		fmt.Fprintf(w, "%s", data)
	}
}

// Service 模拟业务获取数据
func Service(ch chan []byte) {
	if rand.Intn(50) > 10 { // 一定几率获得数据
		ch <- []byte("{\"data\":\"test\"}")
	}
}
