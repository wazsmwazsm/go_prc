package main

import "fmt"

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    // 对接口保存的底层对象为 nil 时的情况进行处理
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

func main()  {
    var i I // 此时 i 为 nil
    var t *T // 一个为 nil 的 T 类型变量

    i = t // 注意，i 保存的底层变量为 nil，但是自身不为 nil
    describe(i)
    i.M()

    i = &T{"hello"}
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
