package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := new(User)
	name := (*string)(unsafe.Pointer(u))
	*name = "test"

	// unsafe.Pointer 转化为 uintptr 来计算，再转换为 unsafe.Pointer
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.Age)))
	*age = 25
	fmt.Println(u)
}
