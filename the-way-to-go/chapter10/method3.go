package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p * Point) Abs() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

type NamePoint struct {
	Point // 内嵌 Point 后，NamePoint 会拥有 Point 的方法
	name string
}

func main() {
	n := &NamePoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())
}
