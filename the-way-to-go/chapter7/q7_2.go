package main

import "fmt"

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	fmt.Printf("%c\n", b[1:4]) // ['o', 'l', 'a']
	fmt.Printf("%c\n", b[:2]) // ['g', 'o']
	fmt.Printf("%c\n", b[2:]) // ['l', 'a', 'n', 'g']
	fmt.Printf("%c\n", b[:]) // ['g', 'o', 'l', 'a', 'n', 'g']
}
