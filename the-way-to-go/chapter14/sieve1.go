package main

import (
	"fmt"
)

func generate(ch chan int) {
	for i := 2; ; i++ { // 不断生产大于 2 的正整数
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 { // 返回所有不能被 prime 整除的值
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)

	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		// 对于每个 prime，生成一个过滤器, 过滤器的入口是上个过滤器的出口
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
