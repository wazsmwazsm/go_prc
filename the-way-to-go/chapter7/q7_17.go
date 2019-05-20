package main

import "fmt"

func main() {
	s := []int{54, 2, 11, 45, 93, 13, 5, 7, 6}
	s = mapFunc(func(i int) int {
		return i * 10
	}, s)
	fmt.Println(s)
}

func mapFunc(fn func (int) int, s []int) []int {
	list := make([]int, len(s))

	for k, v := range s {
		list[k] = fn(v)
	}

	return list
}

