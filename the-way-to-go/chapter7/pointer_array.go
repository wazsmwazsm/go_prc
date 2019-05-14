package main

import "fmt"

func main() {
	var a [3]int
	f(a)
	fp(&a)
}

func f(a [3]int) { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }
