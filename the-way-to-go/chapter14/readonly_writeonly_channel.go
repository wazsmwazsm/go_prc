package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// ch := make(chan<- int) // 可以这么定义，但是单向的 channel 变量完全没意义
	go source(ch)
	go sink(ch)

	time.Sleep(time.Second)
}

func source(ch chan<- int) { // 只写 channel，传入的 channel 只能写数据
	for i := 0; i < 5; i++ {
		ch <- i
		// <-ch // 会报错
	}
}

func sink(ch <-chan int) { // 只读 channel，传入的 channel 只能读数据
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
		// ch <- 1 // 会报错
	}
}
