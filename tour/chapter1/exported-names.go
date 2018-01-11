package main

import (
    "fmt"
    "math"
)

func main() {
    // go 只支持首字母大写的函数变量导出，此处应该为 math.Pi
    fmt.Println(math.pi)
}
