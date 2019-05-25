package main

import "fmt"

func main()  {
    ch := make(chan int, 2)

    // 在 goroutine 中，ch 超出 2 个容量的时候阻塞掉
    // 但是随着主线程中 range 的读取, ch 容量空出来的时候它会继续执行
    // 由于没有阻塞所有线程，所以不是死锁
    go func() {
        for i := 0; i < 5; i++ {
           ch <- i
        }
        close(ch)
    }()

    // 在主 goroutine 中, 这样写是死锁, 因为阻塞了主线程
    // for i := 0; i < 5; i++ {
    //    ch <- i
    // }
    // close(ch)

    for i := range ch {
        fmt.Println(i)
    }

}
