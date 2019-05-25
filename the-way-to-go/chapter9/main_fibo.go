package main

import (
	"fmt"
	"./packtest/fibo"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(fibo.Get(i))
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(fibo.GetByOperator(i, "+"))
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(fibo.GetByOperator(i, "*"))
	}
}