package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		
		time.Sleep(1) // 沉睡 1ns 
	}
	end := time.Now()
	delta := end.Sub(start) // 获取 start 到 end 的时间
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

