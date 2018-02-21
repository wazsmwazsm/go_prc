package main

import (
    "fmt"
    "math"
)

// 穿一个函数当参数
// 声明函数的类型
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    // 匿名函数
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }

    fmt.Println(hypot(5, 12))
    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))

}
