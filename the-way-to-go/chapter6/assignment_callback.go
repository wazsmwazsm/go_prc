package main

import "fmt"

func main() {
	// 匿名函数可以赋值
	fv := func() { fmt.Println("Hello world") }

	fmt.Printf("type: %T, value: %v\n", fv, fv)

	fv()
}
