package main

import "fmt"

func main() {
	x := min(8, 3, 4, 9, 2, 7, 6)
	fmt.Printf("The minimum is: %d\n", x)

	slice := []int{8, 5, 6, 2, 4, 1, 7}
	x = min(slice...) // 结构 slice
	fmt.Printf("The minimum is: %d\n", x)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}

	min := s[0]

	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}
