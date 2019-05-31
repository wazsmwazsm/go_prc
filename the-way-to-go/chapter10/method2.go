package main

import "fmt"

type IntVector []int // 方法的接收者必须和方法在同一个包里，不能给 int 、[]int 这些直接定义方法

func main() {
	fmt.Println(IntVector{1, 2, 3}.Sum())
}

// 非结构体接收者
func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}
