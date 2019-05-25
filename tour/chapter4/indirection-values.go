package main

import (
		"fmt"
		"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main()  {
		v := Vertex{3, 4}
		fmt.Println(v.Abs())
		fmt.Println(AbsFunc(v)) // 必须传值

		p := &Vertex{4, 3}
		fmt.Println(p.Abs())	// 虽然是指针，但是还可调用，被 go 解释为 (*p).Abs()
		fmt.Println(AbsFunc(*p))  // 必须传指针
}
