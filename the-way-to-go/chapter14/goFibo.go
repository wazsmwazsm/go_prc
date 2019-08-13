package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	term := 25
	i := 0
	ch := make(chan int)

	start := time.Now()
	go fiboterms(term, ch)
	for {
		if res, ok := <-ch; ok {
			fmt.Printf("fibonacci(%d) is: %d\n", i, res)
			i++
		} else {
			end := time.Now()
			delta := end.Sub(start)
			fmt.Printf("longCalculation took this amount of time: %s\n", delta)
			os.Exit(0)
		}
	}
}

func fiboterms(term int, ch chan int) {
	for i := 0; i < term; i++ {
		ch <- fibo(i)
	}
	close(ch)
}

func fibo(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibo(n-1) + fibo(n-2)
	}
	return res
}
