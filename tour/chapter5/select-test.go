package main

import "fmt"

func main() {
    // // 主 goroutine 中 select
    // ch1 := make (chan int, 1)
    // ch2 := make (chan int, 1)
    //
    // // ch2 <- 1
    //
    // select {
    //     // 如果所有分支都没有满足解除阻塞条件的也无 default, 主线程就阻塞了
    //     case <-ch1:
    //         fmt.Println("ch1 pop one element")
    //     case <-ch2:
    //         fmt.Println("ch2 pop one element")
    //     // 使用 default select 就不会阻塞
    //     // default:
    //     //   fmt.Println("default")
    // }

    // gonotine 中 select

    ch1 := make (chan int)
    ch2 := make (chan int)

    ct := 0 // 计 for 的循环数

    // 开启一个 goroutine, 用 for 不停监测 select
    go func() {
      // 使用 for 循环不停监测
      for {
          fmt.Println("select begin")

          select {
          case ch1 <- 1:
              fmt.Println("ch1 push one element")
          case <- ch2:
              fmt.Println("ch2 pop one element")
          // default:
          //     fmt.Println("default")
          }
          ct++
          fmt.Println("select end")
      }

    }()

    // 读取 ch1
    df := <- ch1
    fmt.Println("df",df)

    // 开另一个 goroutine 一直和主 goroutine 同步数据, 防止主 goroutine 执行完毕
    ch := make(chan int)

    go func() {
        i := 0
        for {
            i++
            ch <- i
        }
        close(ch)
    }()

    for i := range ch {
        // i = i
        if i == 200 {
            // 尝试执行 select 分支
            ch2 <- i
            ch2 <- i
        }
        // 退出和主 goroutine 的同步
        if i == 50000 {
            fmt.Println("ct", ct) // 查看第一个 goroutine 中 for 循环了几次
            break
        }
    }
}
