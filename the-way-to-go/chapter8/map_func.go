package main

import "fmt"


func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 40 },
		3: func() int { return 70 },
	}

	fmt.Println(mf)
}
