package main

import (
    "fmt"
    "math"
)

// func Sqrt(x float64) float64 {
//     // 定义误差标准
//     eps := 0.0001
//     result := float64(x)
//
//     for true {
//         lastValue := result
//         result = result - (result * result - x) / (2 * result)
//         if math.Abs(result - lastValue) < eps {
//             break
//         }
//     }
//
//     return result
// }

// 误差值 1 的 -6 次方
const delta = 1e-6

func Sqrt(x float64) float64 {
    z := x
    n := 0.0
    for math.Abs(n - z) > delta {
        n, z = z, z - (z * z - x) / (2 * z)
    }

    return z
}

func main() {
    const x = 2
    mine, theirs := Sqrt(x), math.Sqrt(x)
    fmt.Println(mine, theirs, mine - theirs)
}
