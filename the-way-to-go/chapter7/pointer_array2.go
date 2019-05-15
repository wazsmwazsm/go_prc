package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}
}

func fp(a *[3]int) { fmt.Println(a) }
