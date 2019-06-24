package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}

func badCall() {
	a, b := 1, 0
	n := a / b
	fmt.Println(n)
}

func test() {
	defer func() {
		// 捕获除 0 panic
		if r := recover(); r != nil {
			fmt.Printf("Panicing %s\r\n", r)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n")
}
