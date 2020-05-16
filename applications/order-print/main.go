// 三个 go 程，要求按照顺序打印
// 利用 channel 的特性即可。阻塞等待 channel
package main

import (
	"fmt"
	"time"
)

var ch1, ch2 chan bool

func main() {
	ch1, ch2 = make(chan bool), make(chan bool)
	go one()
	go two()
	go three()
	time.Sleep(1e9)
}

func one() {
	fmt.Println("one")
	ch1 <- true
}

func two() {
	<-ch1
	fmt.Println("two")
	ch2 <- true
}
func three() {
	<-ch2
	fmt.Println("three")
}
