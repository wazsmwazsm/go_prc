package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    // 如果是负数，开平方得到的是虚数
    if x < 0 {
        return sqrt(-x) + "i"
    }
    // 转换为 string 输出
    return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
    if v:= math.Pow(x, n); v < lim {
        return v
    }
    // return v 在 if 外是读取不到 v 的
    return lim
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))

    fmt.Println(
  		pow(3, 2, 10),
  		pow(3, 3, 20),
  	)
}
