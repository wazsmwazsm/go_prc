package main

import "fmt"

// 可以指定类型
const Pi float64 = 3.14

const (
    Big   = 1 << 100    // 根据长度自动确定类型
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    const World = "世界"
    fmt.Println("Hello", World)
	  fmt.Println("Happy", Pi, "Day")

    const Truth = true
	  fmt.Println("Go rules?", Truth)

    fmt.Println(needInt(Small))
  	fmt.Println(needFloat(Small))
  	fmt.Println(needFloat(Big))
}
