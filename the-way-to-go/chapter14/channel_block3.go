package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(5 * 1e9) // 5s 后读取 channel

		x := <- ch
		fmt.Println("received", x)
	}()

	fmt.Println("sending", 10)

	ch <- 10
	fmt.Println("sent", 10)
}

