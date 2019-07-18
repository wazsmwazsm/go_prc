package main

import (
	"fmt"
	"./pkg/mysync"
)

var se mysync.Semaphore 

func main() {
	se = make(mysync.Semaphore, 10) // 信号量的缓冲长度最好设置为 wait 个数，最省资源
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer se.Signal()
			fmt.Println(i)
		}(i)
	}
	se.Wait(10) // 等待 10 个信号（小于发送次数会 deadlock）
}

