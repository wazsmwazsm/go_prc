package main

import "fmt"

func main() {
    i, j := 42, 2701

    // & 符号生成一个指针
    p := &i
    // * 可以取指针的值
    fmt.Println(*p)
    // 改变指针指向的值
    *p = 21
    fmt.Println(i) // i 也会改变

    // 将 p 指向 j 的地址
    p = &j
    *p = *p / 37
    fmt.Println(j)
}
