package main

import (
	"fmt"
)

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

	panic("bad test !!")
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
}
