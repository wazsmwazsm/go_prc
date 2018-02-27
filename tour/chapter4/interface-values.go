package main

import (
  	"fmt"
  	"math"
)

type I interface {
    M()
}

type T struct {
    S string
}
// go 的接口是隐性实现
// 只要底层类型变量定义了这个方法就认为次变量实现了接口
func (t *T) M() {
    fmt.Println(t.S)
}

// 另一个类型实现接口
type F float64

func (f F) M() {
    fmt.Println(f)
}

func main()  {
    // 声明接口值
    var i I
    // 保存一个具体底层类型的具体值
    i = &T{"hello"}
    describe(i)
    i.M()

    i = F(math.Pi)
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
