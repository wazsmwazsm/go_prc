package main

import "fmt"

func main() {
    /*
    手动 fallthrough
    和 C 系不同的是
    fallthrough 不会管后面的
    条件是否满足，会强行向下执行
    */
    fmt.Println("Fall through ")
    num := 10
    switch {
        case num < 30:
            fmt.Println("num < 30.")
            fallthrough
        case num < 20:
            fmt.Println("num < 20.")
            fallthrough
        case num < 10:
            fmt.Println("num < 10.")
    }
}
