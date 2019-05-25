package main

import "fmt"
import "strings"

func main() {
	w, h := 20, 10
	// 1
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 2
	str := strings.Repeat("*", w)
	for y := 0; y < h; y++ {
		fmt.Println(str)
	}
}
