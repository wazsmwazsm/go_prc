package workpool

import (
	"testing"
)

func BenchmarkPool(b *testing.B) {
	demo := func() {
		a := 0
		for i := 0; i < 1000000; i++ {
			a = a + i
			b := 0
			b = b + i
		}
	}

	pool := NewPool(50)

	for i := 0; i < b.N; i++ {
		pool.Run(func() {
			demo()
		})
	}
}

func BenchmarkGo(b *testing.B) {
	demo := func() {
		a := 0
		for i := 0; i < 1000000; i++ {
			a = a + i
			b := 0
			b = b + i
		}
	}

	for i := 0; i < b.N; i++ {
		go demo()
	}
}
