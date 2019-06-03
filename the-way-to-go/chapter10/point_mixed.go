package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Polar struct {
	x, y, z float64
}

func Abs(p *Point) float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func Scale(p *Point, s float64) {
	p.x = p.x * s
	p.y = p.y * s
}

func (p *Polar) Abs() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y + p.z * p.z)
}

func (p *Polar) Scale(s float64) {
	p.x = p.x * s
	p.y = p.y * s
	p.z = p.z * s
}

func main() {
	point := Point{2.3, 1.1}
	fmt.Println(Abs(&point))
	Scale(&point, 4)
	fmt.Println(point)
	fmt.Println(Abs(&point))

	polar := &Polar{1.2, 10.4, 6.5}
	fmt.Println(polar.Abs())
	polar.Scale(4)
	fmt.Println(polar)
	fmt.Println(polar.Abs())
}
