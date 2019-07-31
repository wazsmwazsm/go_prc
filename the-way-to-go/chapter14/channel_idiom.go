package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump()
	go suck(stream)

	time.Sleep(time.Second)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 9; i++ {
			ch <- i
		}
	}()

	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
