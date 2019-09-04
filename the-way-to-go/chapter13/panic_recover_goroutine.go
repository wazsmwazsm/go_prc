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
	// 如果产生 panic 的程序以另一个 gorouting 运行，则捕捉不到
	// 因为虽然 badCall 在 test 中启动，但每个 gorouting 都单独运行，
	// 都受控于 main gorouting, 所以 badCall gorouting panic 后在
	//  test 的 defer 是不会走到的，会直接 panic 到 main gorouting，
	// 导致程序崩溃, 所以正确方式应该在启动 gorouting 时做捕捉
	// go badCall()
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
