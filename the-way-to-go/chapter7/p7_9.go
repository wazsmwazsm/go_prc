package main

import "fmt"

func main() {
	s := []int{1,3,4,5}
	fmt.Println(s, len(s), cap(s))
	newSlice := make([]int, len(s) * 3)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	newSlice = make([]int, len(s) * 0)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
}
