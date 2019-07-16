package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10) // 带 buffer

	go func() {
		time.Sleep(5 * 1e9) // 5s 后读取 channel

		x := <- ch
		fmt.Println("received", x)
	}()

	fmt.Println("sending", 10)

	ch <- 10 // channel 没满，发送后不阻塞
	fmt.Println("sent", 10)
}

