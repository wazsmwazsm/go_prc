package main

import (
	"fmt"
)

type Request struct {
	a, b    int
	chreply chan int
}

func (r *Request) String() string {
	// 打印 Request 变量时读取响应
	return fmt.Sprintf("%d+%d=%d", r.a, r.b, <-r.chreply)
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

	// make requests:
	req1 := &Request{3, 4, make(chan int)}
	req2 := &Request{150, 250, make(chan int)}
	// send requests on the service channel
	adder <- req1
	adder <- req2
	// ask for the results: ( method String() is called )
	fmt.Println(req1, req2)
	// shutdown server:
	quit <- true
	fmt.Print("done")
}
