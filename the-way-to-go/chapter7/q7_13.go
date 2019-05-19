package main

import "fmt"

func main() {
	s := "Hello My Friend!"
	fmt.Println(s[len(s)/2:] + s[:len(s)/2])
}

