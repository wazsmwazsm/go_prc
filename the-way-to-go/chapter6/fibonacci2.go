package main

import "fmt"

func main() {
	result, n := 0, 0
	for i := 0; i <= 10; i++ {
		n, result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", n, result)
	} 
}

func fibonacci(n int) (index, res int) {
	index = n
	if n <= 1 {
		res = 1
	} else {
		_, n1 := fibonacci(n - 1)
		_, n2 := fibonacci(n - 2)
		res = n1 + n2
	}

	return
}

