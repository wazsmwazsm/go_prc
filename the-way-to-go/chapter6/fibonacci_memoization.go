// 在递归时使用保存之前计算结果的方式来提升效率，其实使用闭包的话已经能达到这个要求
// 这个只是为了演示，为了递归提升效率去做，实际项目请用闭包方式
package main

import (
	"fmt"
	"time"
)

// 演示中计算 41 位
const LIM = 41
// 用来保存计算结果
var fibs [LIM]uint64

func main() {
	start := time.Now()
	for i := 0; i < LIM; i++ {
		fibonacci(i)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res uint64) {
	// 如果该结果已经被缓存，则无需计算
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	// 保存计算结果
	fibs[n] = res
	return
}
