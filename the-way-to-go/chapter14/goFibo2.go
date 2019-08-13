package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 25)
	start := time.Now()
	go fibo(cap(ch), ch)
	for v := range ch {
		fmt.Println(v)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}

func fibo(n int, ch chan<- int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, y+x
	}
	close(ch)
}
