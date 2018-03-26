package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int() // 返回 0 - 1 之间的随机数, 包括 0 不包括 1
		fmt.Printf("%d / ", a)
	}

	for i := 0; i < 5; i++ {
		r := rand.Intn(8) // 函数 rand.Intn 返回介于 [0, n) 之间的伪随机数。
		fmt.Printf("%d / ", r)
	}

	fmt.Println()
	// 生成随机种子
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens) // 初始化种子
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100 * rand.Float32())
	}
}
