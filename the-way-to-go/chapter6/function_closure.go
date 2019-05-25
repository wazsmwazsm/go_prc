package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ") // 1
	fmt.Print(f(20), " - ") // 21
	fmt.Print(f(300)) // 321
}

func Adder() func(int) int {
	// 这里的 x 充当一个 static 变量的角色
	// 绑定在闭包上不会销毁，每次调用 x 都会保留之前的值
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
