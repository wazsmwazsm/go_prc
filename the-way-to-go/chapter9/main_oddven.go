package main 

import (
	"fmt"
	"./packtest/even"
)

func main() {
	for i := 1; i <= 100; i++ {
		if even.IsEven(i) {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}

}