package main

import "fmt"

type Rectangle struct {
	width, height float64
}

// 函数方式
// 通过传入类型实例操作
func area(r Rectangle) float64 {
	return r.width*r.height
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))
}
