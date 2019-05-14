package main

import "fmt"

func main() {
	a := [5]int{1, 2, 4, 5, 7}
	f(a)
	fmt.Printf("Outside func, a = %v\n", a)

	b := a // 数组是值类型，赋值会拷贝数据
	b[2] = 111
	fmt.Printf("b = %v\n", b)
	fmt.Printf("a = %v\n", a)

}

func f(a [5]int) { // 数组是值类型，传入函数是数据拷贝
	a[0] = 12
	fmt.Printf("Inside func, a = %v\n", a)
}