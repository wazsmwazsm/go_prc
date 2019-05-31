package main

import (
	"fmt"
	"unsafe"
)

type Foo map[string]string
type Bar struct {
	one string
	two int
}
func main() {
	// ok
	x := Bar{"hello,hahahahaha!", 23}
	fmt.Println(unsafe.Sizeof(x)) // 查看结构体占用的内存

	y := new(Bar)
	y.one = "hello,hahahahaha!"
	y.two = 23
	fmt.Println(unsafe.Sizeof(y)) // 这里 y 是指针, 只占用一个字 (64 位系统 8 字节)

	// not ok
	// z := make(Bar) // 编译错误

	// ok 
	m := make(Foo)
	m["x"] = "hello"
	m["y"] = "world"

	// not ok
	// u := new(Foo)
	// u["x"] = "a" // 会报错
}
