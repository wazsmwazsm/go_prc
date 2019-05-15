package main

import "fmt"

func main() {
	s1 := []float32{2.33, 9.2, 9.2}
	fmt.Println(sum(s1))

	s2 := []float32{11.2, 23.4, 99.2, 1001.1, 1982.2229}
	fmt.Println(sum(s2))
}

func sum(s []float32) (sum float32) {
	for _, v := range s {
		sum += v
	}
	return 
}
