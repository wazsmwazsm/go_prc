package main

import "fmt"

func main() {
    pow := make([]int, 10)

    // 只取得索引
    for i := range pow {
        // 将索引左移一位 ( 2 的次方 + 1 )
        pow[i] = 1 << uint(i)
    }
    // 通过赋值给 _ 来忽略序号和值
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
}
