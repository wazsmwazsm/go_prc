package main 

import "fmt"

func f1(in chan int) {
	fmt.Println(<- in)
}

func main() {
	out := make(chan int)
	// go f1(out) // 不死锁
	out <- 2 // 死锁, f1 的 go 程还没有启动，main go 程先阻塞读取 channel 了
	go f1(out)
}
