package main

import (
	"fmt"
)

type Request struct {
	a, b    int
	chreply chan int
}

type binOp func(a, b int) int

// 运行 handler 返回响应
func run(op binOp, req *Request) {
	req.chreply <- op(req.a, req.b)
}

// server
func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service: // 有请求来，则开启一个 goroutine 来处理
			go run(op, req)
		case <-quit: // 退出服务
			break
		}
	}
}

// 开启 server，返回 request API
func startServer(op binOp) (chan *Request, chan bool) {
	reqChan := make(chan *Request)
	quit := make(chan bool)
	go server(op, reqChan, quit)
	return reqChan, quit
}

func main() {
	// server 开启，设置请求处理 handle
	adder, quit := startServer(func(a, b int) int {
		return a + b
	})

	// 请求模拟
	const count = 100
	var reqs [count]Request
	for i := 0; i < count; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + count
		// 建立 response 接口
		req.chreply = make(chan int)
		adder <- req // 模拟一个请求
	}

	// 读取响应
	for i := count - 1; i >= 0; i-- {
		if <-reqs[i].chreply != count+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}
	quit <- true
	fmt.Println("done")
}
