package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go tel(ch, quit)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-quit:
			os.Exit(0)
		}
	}
}

func tel(ch chan<- int, quit chan<- bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}
