package main

import "fmt"

func main() {
	fibonacci()
}

func fibonacci() {
	var a [50]int

	for i := 0; i < len(a); i++ {
		if i == 0 || i == 1 {
			a[i] = 1
		} else {
			a[i] = a[i - 1] + a[i - 2]
		}
	}

	fmt.Println(a)
}