package main

import "fmt"

func main() {
	s := []int{4, 9, 99, 3, 2, 33, 8}

	fmt.Println(minSlice(s))
	fmt.Println(maxSlice(s))
}

func minSlice(s []int) int {
	tmp := s[0]
	for _, v := range s {
		if v < tmp {
			tmp = v
		}
	}

	return tmp
}

func maxSlice(s []int) int {
	tmp := s[0]
	for _, v := range s {
		if v > tmp {
			tmp = v
		}
	}

	return tmp
}
