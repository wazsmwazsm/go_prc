package main

import (
	"fmt"
	"time"
)

var num int = 0
var ch chan int

func main() {
	ch = make(chan int)

	for i := 0; i < 1000; i++ {
		go add(ch)
	}
	// channel 的数据同步方式
	for i := 0; i < 1000; i++ {
		num += <- ch
	}
	fmt.Println(num)
}

func add(ch chan int) {
	time.Sleep(1e9) // 模拟耗时请求
	ch <- 1
}
