package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main()  {
    v := Vertex{3, 4}
    v.Scale(2)  // 虽然是指针接收者，但是值任然可调用, 会被解释为 (&p).Scale(5)
    ScaleFunc(&v, 10) // 必须传指针

    p := &Vertex{4, 3}
    p.Scale(3)
    ScaleFunc(p, 8) // 必须传值

    fmt.Println(v, p)
}
