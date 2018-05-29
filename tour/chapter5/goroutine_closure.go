package main

import (
	"fmt"
	"time"
)

// 多个 goroutine 共享闭包绑定的参数时的问题

func main() {

	// 这样打印的可能是 5 5 5 5 5 (5 个闭包都绑定了 i 变量，而 i 最后的值是 5，编译器直接使用 5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(1)

	// 解决方案 1，传入参数到闭包
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(1)

	// 解决方案 2，新建变量即可 (这样每个闭包绑定的就不是同一个变量了)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(1)
}
