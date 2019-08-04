package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 3*time.Second) // 超时会给 Done 的 channel 发送结束信号

	go work(ctx)

	time.Sleep(10 * time.Second) // 给 gorutine 时间执行
}

func work(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done(): // 调用 cancel handler 后，channel 可读，解除阻塞
			fmt.Println("done")
			return
		default:
			if deadline, ok := ctx.Deadline(); ok {
				fmt.Println("deadline:", deadline)
			}
			fmt.Println("work")
		}
	}
}
