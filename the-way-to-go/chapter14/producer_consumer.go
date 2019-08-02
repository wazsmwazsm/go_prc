package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go producer(ch)
	go comsumer(ch, done)
	<-done
}

func producer(in chan int) {
	for i := 0; i < 90; i += 10 {
		in <- i
	}
	close(in)
}

func comsumer(out chan int, done chan bool) {
	for v := range out {
		fmt.Println(v)
	}
	done <- true
}
