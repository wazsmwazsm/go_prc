package main

import "fmt"

func main() {
	sum, prod, diff := twoInt(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
	sum, prod, diff = twoInt2(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
}

func twoInt(x, y int) (int, int, int) {
	return x + y, x * y, x - y
}

func twoInt2(x, y int) (a, b, c int) {
	a, b, c = x + y, x * y, x - y
	return
}
