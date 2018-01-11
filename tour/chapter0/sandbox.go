package main

// 打包导入，可以写成
// import "fmt"
// import "time"
import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Welcome to the playground!")
    fmt.Println("The time is", time.Now())
}
