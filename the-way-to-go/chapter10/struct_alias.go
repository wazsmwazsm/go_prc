package main

import "fmt"

type A struct {
	ax, ay int
}

type B A

func main() {
	b := B{2, 4}
	fmt.Println(b.ax, b.ay)
	fmt.Println(b.sum()) // 虽然 B 是 A 的别名类，但是方法不会共用
}

// 方法是区分接受者的
func (a *A) sum() int {
	return a.ax + a.ay
}
