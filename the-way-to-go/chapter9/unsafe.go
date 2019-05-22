package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 23
	fmt.Println(unsafe.Sizeof(a)) // 返回字节长度 64 bit 系统 int 为 4 字节
}
