
package main

import "fmt"

func fibonacci() func() int {
    f, g := 1, 0
    return func() int {
        // 赋值之前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值
        f, g = g, f + g
        return g
    }
}

func main()  {
    f := fibonacci()

    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
