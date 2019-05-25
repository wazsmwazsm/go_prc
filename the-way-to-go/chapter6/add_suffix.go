package main

import (
	"fmt"
	"strings"
)

func main() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	fmt.Println(addBmp("b12.bmp"))
	fmt.Println(addBmp("b12"))
	fmt.Println(addJpeg("b13.jpeg"))
	fmt.Println(addJpeg("b13"))
	fmt.Println(addJpeg("b12.txt"))
}

// 生产同一个模式的函数，工厂函数
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
