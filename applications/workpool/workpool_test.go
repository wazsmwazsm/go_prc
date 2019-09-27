package workpool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestRelease(t *testing.T) {
	wg := new(sync.WaitGroup)
	demo := func() {
		for i := 0; i < 1000000; i++ {
		}
	}

	task := func() {
		wg.Add(1)
		defer wg.Done()
		demo()
	}

	pool := NewPool(20)
	for i := 0; i < 100; i++ {
		pool.Run(task)
	}
	fmt.Printf("%v\n", runtime.NumGoroutine())
	pool.Release()
	time.Sleep(time.Second * 2)
	fmt.Printf("%v\n", runtime.NumGoroutine())
}

func BenchmarkPool(b *testing.B) {
	// 使用协程池 bench 时内存明显要少很多
	demo := func() {
		for i := 0; i < 1000000; i++ {
		}
	}

	task := func() {
		demo()
	}

	pool := NewPool(20)
	for i := 0; i < b.N; i++ {
		pool.Run(task)
	}
}

func BenchmarkPoolFilled(b *testing.B) {
	// 使用协程池 bench 时内存明显要少很多
	demo := func() {
		for i := 0; i < 1000000; i++ {
		}
	}

	task := func() {
		demo()
	}

	pool := NewPool(20)
	pool.Fill()
	for i := 0; i < b.N; i++ {
		pool.Run(task)
	}
}

func BenchmarkGo(b *testing.B) {
	demo := func() {
		for i := 0; i < 1000000; i++ {
		}
	}

	for i := 0; i < b.N; i++ {
		go demo()
	}
}
