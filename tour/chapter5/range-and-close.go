package main

import "fmt"

func main() {
    c := make(chan int, 10)
    // 在 goroutine 中，ch 超出 buffer 容量的时候阻塞掉
    // 但是随着主 goroutine 中 range 的读取, ch 容量空出来的时候它会继续执行
    // 由于没有阻塞所有 goroutine ，所以不是死锁
    // 当然，这样会频繁的切换 goroutine, 不建议这么做
    go fibonacci(15, c)
    // 循环取出 channel 中数据
    for i := range c {
        fmt.Println(i)
    }
}

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        // 数据存入 channel
        c <- x
        x, y = y, x + y
    }
    // 关闭 channel
    // 可以在主 goroutine 中关闭
    // 不关闭的话 range 读取需要判断, 不然会死锁
    close(c)
}
