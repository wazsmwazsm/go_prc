package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx) // 返回子节点和 cancel handler

	go work(ctx)

	time.Sleep(5 * time.Second)
	cancel()                    // 向 Done 的 channel 发送结束信号
	time.Sleep(1 * time.Second) // 给 goroutine 留执行时间
}

func work(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done(): // 调用 cancel handler 后，channel 可读，解除阻塞
			fmt.Println("done")
			return
		default:
			fmt.Println("work")
		}
	}
}
