package main

import (
	"fmt"
)

func main() {
	quit := make(chan int)
	
	close(quit) // close 可以解除进程的阻塞

	q := <- quit

	fmt.Println(q) // 0
}