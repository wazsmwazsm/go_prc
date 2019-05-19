package main

import "fmt"

func main() {
	s := []int{54, 2, 11, 45, 93, 13, 5, 7, 6}
	bubbleSort(s)
	fmt.Println(s)
}

func bubbleSort(s []int) {
	var tmp int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s) - 1 - i; j++ {
			if s[j] > s[j + 1] {
				tmp = s[j]
				s[j] = s[j + 1]
				s[j + 1] = tmp
			}
		}
	}
}

