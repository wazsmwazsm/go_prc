// 请求限流
package main

import (
	"fmt"
	"sync"
	"time"
)

const MAXREQS = 50

var sem = make(chan int, MAXREQS)
var mu = new(sync.Mutex)
var limitCount = 0

type Request struct {
	a, b    int
	chReply chan int
}

func process(r *Request) {
	// do something
	time.Sleep(1e4) // 睡眠一段时间，让请求可以被限流
}

func handler(r *Request) {
	if len(sem) == MAXREQS { // 统计限流次数
		countLimits()
	}
	sem <- 1 // 向通道写值，如果通道长度等于 MAXREQS 会阻塞，起到限流作用
	process(r)
	<-sem // 执行完毕，流量限制 - 1
}

func server(service chan *Request) {
	for {
		req := <-service
		go handler(req)
	}
}

func countLimits() {
	mu.Lock()
	limitCount++
	fmt.Printf("限流 %v 次\n", limitCount)
	mu.Unlock()
}

func main() {
	service := make(chan *Request)
	go server(service)

	for i := 0; i < 1000; i++ {
		go func() {
			chReply := make(chan int)
			service <- &Request{i, i + 100, chReply}
		}()
	}
}
