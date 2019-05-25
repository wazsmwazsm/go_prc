package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 45, 6, 7, 8, 9, 3, 44, 23, 23}

	s = filter(s, func(v int) bool {
		if v % 2 == 0 {
			return true
		}
		return false
	})

	fmt.Println(s)
}

func filter(s []int, fn func(int) bool) []int {
	var newSlice []int
	
	for _, v := range s {
		if fn(v) {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}
