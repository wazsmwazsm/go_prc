package main

import "fmt"

func main() {
	n := 0
	// 这里只是举例
	// 当引用一个内存占用大的变量时使用指针的优势会更明显
	// 节省了一次内存复制的开销
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply)
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}
