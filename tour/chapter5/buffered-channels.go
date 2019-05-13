package main

import "fmt"

func main()  {
    ch := make(chan int, 2)
    // 仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
    // 非缓存区信道每次发后就阻塞等待读取，会频繁的信道切换，从这点看，buffer 信道某些场景下性能更好
    ch <- 1
    ch <- 2
    // ch <- 3 // 死锁
    fmt.Println(<-ch)
      fmt.Println(<-ch)
    // close(ch) // 这里关闭后，下面语句就不会死锁
    fmt.Println(<-ch) // 缓冲区为空时, 阻塞, 死锁
}
