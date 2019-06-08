package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * (c.radius * c.radius)
}

type Triangle struct {
	bottom, height float32
}

func (t Triangle) Area() float32 {
	return 0.5 * (t.bottom * t.height)
}

func main() {
	r := Rectangle{5, 3} // Rectangle 的值变量实现了 Shaper 接口
	q := &Square{5}  // Square 的指针变量实现了  Shaper 接口
	c := Circle{5}  
	t := Triangle{6., 7.5}

	shapes := []Shaper{r, q, c, t}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
