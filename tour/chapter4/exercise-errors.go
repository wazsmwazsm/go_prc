package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    // 注意不要直接打印 e, 要转换一下
    // 因为 ErrNegativeSqrt 类型的变量有 Error 方法
    // 打印 e 时会调用 Error, 在 Error 中打印会无限自调用耗尽内存
    return fmt.Sprintf("%v 是负数！无法开方！", float64(e))
}

const delta = 1e-6

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }

    z := x
    n := 0.0
    for math.Abs(n - z) > delta {
        n, z = z, z - (z * z - x) / (2 * z)
    }

    return z, nil
}

func main()  {

    v, err := Sqrt(2)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(v)
    }

    v, err = Sqrt(-2)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(v)
    }
}
