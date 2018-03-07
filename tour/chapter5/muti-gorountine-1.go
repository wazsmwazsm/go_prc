package main

import "fmt"

func main()  {
    count := 50
    quit := make(chan int)

    // 开启多个 goroutine
    for i := 0; i < count; i++ {
        go foo(i, quit)
    }
    // 结束多个 goroutine
    for i := 0; i < count; i++ {
        <- quit
    }
}

func foo(id int, c chan int) {
    fmt.Println(id)
    c <- 0
}
