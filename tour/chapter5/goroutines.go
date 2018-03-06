package main

import (
  	"fmt"
  	"time"
)

func main()  {
    // 启动一个新的 goroutine
    go say("world")
    // 当前 goroutine 中的程序
    say("hello")
}

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}
