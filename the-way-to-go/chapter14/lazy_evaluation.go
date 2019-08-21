// yield 惰性生成器模式, 每次读取时才生成
package main

import (
	"fmt"
)

var resume chan int

func main() {
	resume = integers()
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
}

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()

	return yield
}

func generateInteger() int {
	return <-resume
}
