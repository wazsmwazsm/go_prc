package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i *int

	p := 10
	i = &p
	// m := (*float64)(i)                 // 不能转换
	m := (*float64)(unsafe.Pointer(i)) // 可以转换
	fmt.Println(m)
}
