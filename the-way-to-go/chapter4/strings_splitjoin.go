package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str) // 用一个多个空白字符将字符串生成 slice
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, v := range sl {
		fmt.Printf("%s - ", v)
	}

	fmt.Println()

	str2 := "GO1|The ABC of Go|25"
	// 用 | 将字符串生成 slice
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, v := range sl2 {
		fmt.Printf("%s - ", v)
	}
	fmt.Println()
	fmt.Println(strings.Join(sl2, "**"))
}
