package main

import "fmt"

func f() (ret int) {
	defer func() {
		// defer 在 return 之后执行，本来命名返回值 ret 为 1，++ 后变 2，然后结束函数
		ret++
	}()
	return 1
}

func main() {
	// 这里返回的是 2
	fmt.Println(f())
}
