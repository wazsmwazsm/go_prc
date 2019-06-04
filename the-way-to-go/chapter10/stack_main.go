package main

import (
	"fmt"
	"./packages/stack"
)

func main() {
	s := stack.NewStack()

	for i := 0; i < stack.LIMIT + 1; i++ {
		if err := s.Push(i + 3); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(s)
		}
	}

	for i := 0; i < stack.LIMIT + 1; i++ {
		if _, err := s.Pop(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(s)
		}
	}
}
