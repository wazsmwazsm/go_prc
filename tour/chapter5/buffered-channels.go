package main

import "fmt"

func main()  {
    ch := make(chan int, 2)
    // 仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
    ch <- 1
    ch <- 2
    // ch <- 3 // 死锁
    fmt.Println(<-ch)
	  fmt.Println(<-ch)
    fmt.Println(<-ch) // 缓冲区为空时, 阻塞, 死锁
}
