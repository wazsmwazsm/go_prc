package main

import "fmt"

func main()  {
    // 空接口可保存任何类型的值
    var i interface {}
    describe(i)

    i = 42
    describe(i)

    i = "hello"
    describe(i)
}

func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}
