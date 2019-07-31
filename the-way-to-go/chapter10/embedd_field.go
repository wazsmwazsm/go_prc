package main

import "fmt"

type MyInt int

type B struct {
	MyInt
	string
	bx, by float32
}

func main() {
	b := B{bx: 5.0, by: 2.2}
	// 匿名字段可以用类型作为名称直接设置、访问
	b.string = "Hello"
	b.MyInt = 13
	fmt.Println(b, b.string, b.MyInt)
}
