package main

import "fmt"

func main() {
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Printf("%c\n", s2) // ['e', 'm']

	s2[1] = 't'
	fmt.Printf("%c %c\n", s1, s2) // ['p', 'o', 'e', 't'] ['e', 't']
}
