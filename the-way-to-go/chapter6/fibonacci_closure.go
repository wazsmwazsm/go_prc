package main

import "fmt"

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a + b
		return b
	}
}
