package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch := make(chan int, 25)
	quit := make(chan bool)
	start := time.Now()
	go fibo(cap(ch), ch, quit)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-quit:
			end := time.Now()
			delta := end.Sub(start)
			fmt.Println(delta)
			os.Exit(0)
		}

	}

}

func fibo(n int, ch chan<- int, quit chan<- bool) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, y+x
	}
	quit <- true
}
