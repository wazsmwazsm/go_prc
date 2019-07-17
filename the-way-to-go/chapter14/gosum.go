package main

import "fmt"

func main() {
	ch := make(chan int)
	go sum(4, 5, ch)
	fmt.Println(<-ch)
}

func sum(x, y int, ch chan int) {
	ch <- x + y
}
