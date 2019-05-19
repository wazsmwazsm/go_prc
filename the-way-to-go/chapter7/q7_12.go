package main

import "fmt"

func main() {
	s := "Hello My Friend!"
	s1, s2 := split(s, 4)
	fmt.Println(s1)
	fmt.Println(s2)
}

func split(s string, i int) (string, string) {
	return s[:i], s[i:]
}