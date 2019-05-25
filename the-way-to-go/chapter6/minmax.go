package main

import "fmt"

func main() {
	var min, max int
	min, max = MinMax(78, 66)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)	
}

func MinMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}

	return b, a
}
