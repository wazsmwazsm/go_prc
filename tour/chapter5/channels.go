package main

import "fmt"

func main()  {
    s := []int{7, 2, 8, -9, 4, 0}
    // 创建一个 channel
    c := make(chan int)
    // 开启两个 goroutine
    // 两个并发执行的函数之间进行同步
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)

    // 注意
    // 在非 buffer channel 中要一个 goroutine 中发数据，另一个 goroutine 中取数据
    // 不然 channel 会一直阻塞等待
    // 默认情况下，发送和接收操作在另一端准备好之前都会阻塞。
    x, y := <-c, <-c // 从 c 中接受
    // x = <-c 死锁
    fmt.Println(x, y, x + y)
}

func sum(s []int, c chan int)  {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // 将和送入 c 中
}
