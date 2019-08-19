// 利用多核版本。效率并不高，只是演示 channel 的使用
package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

const NCPU = 2

func main() {
	runtime.GOMAXPROCS(NCPU)
	start := time.Now()
	fmt.Println(CalculatePi(5000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func CalculatePi(end int) float64 {
	ch := make(chan float64)
	for i := 0; i < NCPU; i++ {
		go term(ch, i*end/NCPU, (i+1)*end/NCPU)
	}
	result := 0.0
	for i := 0; i < NCPU; i++ {
		result += <-ch
	}
	return result
}

func term(ch chan float64, start, end int) {
	result := 0.0
	for i := start; i < end; i++ {
		x := float64(i)
		result += 4 * math.Pow(-1, x) / (2*x + 1)
	}
	ch <- result
}
