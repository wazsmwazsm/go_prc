package main

import "fmt"

func main() {

	c := make(chan int)

	// consumer:
	go func() {
		for {
			fmt.Print(<-c, " ")
		}
	}()

	// producer:
	for {
		select { // 利用 select 的随机选择来随机生成 0 1
		case c <- 1:
		case c <- 0:
		}
	}
}
