package main

import "fmt"

func main() {
    var a []int
    printSlice("a", a)

    // 向 nil slice 追加元素.
    a = append(a, 0)
    printSlice("a", a)

    // 向非空 slice 追加元素.
  	a = append(a, 1)
  	printSlice("a", a)

    // 可以同时追加多个元素
    a = append(a, 2, 3, 4)
  	printSlice("a", a)
}

func printSlice(s string, x []int) {
	 fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
