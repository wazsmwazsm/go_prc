package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()

	// go badCall() // 如果产生 panic 的程序以另一个 gorouting 运行，则捕捉不到
	// 正确的方式
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("Panicing %s\r\n", e)
			}
		}()
		badCall()
	}()

	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
}
