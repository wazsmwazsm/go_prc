package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Second*2)) // 到时间会关闭 done channel

	go work(ctx)

	time.Sleep(10 * time.Second) // 给 gorutine 时间执行
}

func work(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done(): // 调用 cancel handler 后，channel 被关闭，可读，解除阻塞
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
