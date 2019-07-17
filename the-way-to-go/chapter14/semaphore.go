package main

import (
	"fmt"
	"time"
	"./pkg/mysync"
)

var num int = 0

func main() {

	for i := 0; i < 1000; i++ {
		go add()
	}
	
	time.Sleep(1e9)
	fmt.Println(num)
}

func add() {
	// mysync.Mutex 是一个 channel, 长度为 1, Lock 往 channel 写一个数据
	// Unlock 从 channel 取一个数据, 当一个 goroutine 正在执行时, 其它 goroutine
	// 再 Lock 时就会阻塞（获取锁）, 直到这个 goroutine 执行 Unlock, 其它 goroutine
	// 才有机会执行 Lock, 这样就完成了数据的同步, 实现了互斥锁
	mysync.Mutex.Lock()
	defer mysync.Mutex.Unlock()
	num++
}
