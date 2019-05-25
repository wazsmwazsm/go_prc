package main

import (
  	"fmt"
)

type I interface {
    M()
}

type My struct {
    name string
    age int
}

func (m * My) M() {
    fmt.Printf("name %s age %v", m.name, m.age)
}

func main()  {
    m := &My{"Jack", 23}
    i := I(m) // My 实现了接口 I, 可以强转
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
