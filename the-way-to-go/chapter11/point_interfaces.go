package main

import (
	"fmt"
	"math"
)

type Magnitude interface {
	Abs() float64
}

type Point struct {
	x, y float64
}

type Point3 struct {
	x, y, z float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(s float64) {
	p.x = p.x * s
	p.y = p.y * s
}

func (p *Point3) Abs() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y + p.z * p.z)
}

func (p *Point3) Scale(s float64) {
	p.x = p.x * s
	p.y = p.y * s
	p.z = p.z * s
}

func main() {
	var i Magnitude
	point := &Point{2.3, 1.1}
	i = point
	fmt.Println(i.Abs())
	

	point3 := &Point3{1.2, 10.4, 6.5}
	i = point3
	fmt.Println(i.Abs())
	
}
