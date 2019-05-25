package main

import "fmt"

func main()  {
    count := 50
    quit := make(chan int, 50)
    // buffer channel 的话不用开多个 goroutine 就可实现
    for i := 0; i < count; i++ {
        foo(i, quit)
    }

    for i := 0; i < count; i++ {
        <- quit
    }
}

func foo(id int, c chan int) {
    fmt.Println(id)
    c <- 0
}
