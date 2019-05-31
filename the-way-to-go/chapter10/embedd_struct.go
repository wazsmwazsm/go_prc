package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A // 内嵌结构体 (B 会拥有 A 的成员)（如果 B 有和 A 同名的成员，使用 B 的）
	bx, by float32
}

func main() {
	b := B{A{1, 3}, 5.0, 2.2}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}
