package main

import "fmt"

type Integer int

func (p Integer) get() int { 
	return int(p)
}

func f(i int) int {
	return 2 * i
}

func main() {
	var i Integer = 1

	fmt.Println(i)

	// f(i) // 报错
	fmt.Println(f(int(i)))
}
