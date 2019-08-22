// 链式协程 pipeline
package main

import (
	"flag"
	"fmt"
)

var n = flag.Int("n", 10000, "how many goroutines")

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *n; i++ {
		// 将一堆协程用 channel 串起来
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0      // 往 pipe 中传数据, 开始链式反应
	x := <-leftmost // 等待完成
	fmt.Println(x)
}
