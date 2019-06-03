package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

type NamePoint struct {
	Point // 内嵌 Point 后，NamePoint 会拥有 Point 的方法
	name string
}

// 同名时，外层方法会覆盖覆写内嵌类型的方法
func (p *NamePoint) Abs() float64 {
	return p.Point.Abs() * 100.
}

func main() {
	n := &NamePoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())
}
