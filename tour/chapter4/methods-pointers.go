package main

import (
    "fmt"
    "math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
// 为指针接收者声明方法可以修改接收者的值
// 去掉 * 修改会无效，使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
  	v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(10)
    fmt.Println(v.Abs()) // 调用方法
}
