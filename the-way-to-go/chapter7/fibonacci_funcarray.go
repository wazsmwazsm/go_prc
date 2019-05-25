package main

import "fmt"

func main() {
	fmt.Println(fibonacci(20))
}

func fibonacci(limit int) []int {
	s := make([]int, limit)

	for i := 0; i < limit; i++ {
		if i == 0 || i == 1 {
			s[i] = 1
		} else {
			s[i] = s[i - 1] + s[i - 2]
		}
	}

	return s
}

