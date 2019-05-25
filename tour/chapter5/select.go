package main

import "fmt"

func fibonacci(c, quit chan int) {
  	x, y := 0, 1
  	for {
    		select {
			// case 为发送者时, 每次执行 select 都会执行此分支，在数据未被读取时 select 阻塞
    		case c <- x:
    			  x, y = y, x + y
    		case <- quit:
      			fmt.Println("quit")
            // 如果要 break for 的话, 使用 break 标签 到外层
      			return
    		}
  	}
}

func main() {
  	c := make(chan int)
  	quit := make(chan int)
  	go func() {
    		for i := 0; i < 10; i++ {
    			  fmt.Println(<- c)
    		}
    		quit <- 0
  	}()
  	fibonacci(c, quit)
}
