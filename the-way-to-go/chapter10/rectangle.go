package main

import "fmt"

type Rectangle struct {
	width, length int
}

func (r *Rectangle) Area() int {
	return r.width * r.length
}

func (r *Rectangle) Perimeter() int {
	return (r.width + r.length) * 2
}

func main() {
	rectangle := Rectangle{3, 5}
	fmt.Println(rectangle.Area())
	fmt.Println(rectangle.Perimeter())
}
