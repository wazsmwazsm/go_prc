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
        // 为了避免主 goroutine 先执行完毕, 使用 sleep 让主 goroutine 多运行一段时间
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}
