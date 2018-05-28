package main

import "fmt"

func main() {
	// 获取 10 的阶乘
	// 注意 int 型有上限，大数需要 big 包解决
	fmt.Println(factorial(10))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n - 1)
}
