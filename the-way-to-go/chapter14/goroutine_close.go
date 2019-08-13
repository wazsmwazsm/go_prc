package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go tel(ch)
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(i)
	}
}

func tel(ch chan<- int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
}
