package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

// 定义结构体 Vertex 的方法
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs()) // 调用方法
}
