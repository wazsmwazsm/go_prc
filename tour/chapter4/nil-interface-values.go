package main

import "fmt"

type I interface {
    M()
}

func main()  {
    var i I
    // nil 接口值既不保存值也不保存具体类型。
    describe(i)
    // 为 nil 接口调用方法会产生运行时错误
    i.M()

}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
