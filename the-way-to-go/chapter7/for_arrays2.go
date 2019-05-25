package main

import "fmt"

func main() {
	// 编译器计算数组长度
	arr := [...]string{"a", "b", "come", "Hi!"}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Array at index %d is %s\n", i, arr[i])
	}
}
