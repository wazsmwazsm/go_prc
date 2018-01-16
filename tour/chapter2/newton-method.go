package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    // 定义误差标准
    eps := 0.0001
    result := float64(x)

    for true {
        lastValue := result
        result = result - (result * result - x) / (2 * result)
        if math.Abs(result - lastValue) < eps {
            break
        }
    }

    return result
}

func main() {
    fmt.Println(Sqrt(2))
}
