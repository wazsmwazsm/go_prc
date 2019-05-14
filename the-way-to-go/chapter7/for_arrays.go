package main

import "fmt"

func main() {
	var arr [5]int

	for i := 0; i < len(arr); i++ {
		arr[i] = 2 * i
	}

	for k, v := range arr {
		fmt.Printf("Array at index %d is %d\n", k, v)
	}
}
