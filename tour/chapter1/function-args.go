package main

import "fmt"

func myfunc1(args ...int) {
    fmt.Println("myfunc1() invoked")

    for _, arg := range args {
        fmt.Print(arg, " ")
    }

    fmt.Println()
}

func myfunc2(args ...int) {
    fmt.Println("myfunc2() invoked")
    // 将参数传给 myfunc1
    myfunc1(args...)
}

func main() {
    myfunc2(1, 2, 3, 4)
}
