package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Println(a, b)
	c, b := 4, 5 // 只要有一个未定义, := 就可以用
	fmt.Println(a, b, c)
}
