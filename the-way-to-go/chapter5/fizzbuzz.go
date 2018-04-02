package main

import "fmt"

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

func main() {
	for i := 0; i <= 100; i++ {
		switch {
		// 注意这里 FIZZBUZZ 要放在前面, 不然 FIZZ 和 BUZZ 会先满足条件
		// 而结束 switch
		case i % FIZZBUZZ == 0:
			fmt.Println("FizzBuzz")
		case i % FIZZ == 0:
			fmt.Println("FIZZ")
		case i % BUZZ == 0:
			fmt.Println("BUZZ")
		default:
			fmt.Println(i)
		}
	}
}
