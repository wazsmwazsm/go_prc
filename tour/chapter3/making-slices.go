package main

import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a)
    b := make([]int, 0, 5)
	  printSlice("b", b)
    c := b[:2]
  	printSlice("c", c)
    // cap(d) 为 3 请参考这里
    // https://studygolang.com/articles/4879
    // 由于 slice 进行切片后引用的任然是一个数组
    // 引用首指针后移，容量变小 b[:2] 这样的容量不变
  	d := c[2:5]
  	printSlice("d", d)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
