package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	res := fibo()
	for i := 0; i < 25; i++ {
		fmt.Println(<-res)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func dup3(in chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, out := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			res := <-in
			a <- res
			b <- res
			out <- res
		}
	}()

	return a, b, out
}

func fibo() <-chan int {
	in := make(chan int)
	a, b, out := dup3(in)
	go func() {
		in <- 0
		in <- 1
		<-a // 去掉 a channel 的第一个数据，构造 n - 1, n -2 的模式
		for {
			in <- <-a + <-b
		}
	}()

	return out
}
