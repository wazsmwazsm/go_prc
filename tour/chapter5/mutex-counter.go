package main

import (
	"fmt"
	"sync"
	"time"
)
// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v  map[string]int
	mux sync.Mutex
}
// Inc 增加给定 key 的计数器的值
func (c *SafeCounter) Inc(key string) {
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	// 使用 go run -race source.go 来判断是否有数据竞争
	c.mux.Lock()
	// 操作数据
	c.v[key]++
	// 解锁
	c.mux.Unlock()
}
// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	// 保证互斥锁一定会被解锁
	defer c.mux.Unlock()

	c.mux.Lock()

	return c.v[key]
}

func main() {
	c := SafeCounter{
		v: make(map[string]int),
	}

	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	// 结果为 1000, 说明数据安全的操作了, 没有因为并发出错
	fmt.Println(c.Value("somekey"))
}
