package main

import "fmt"

func main() {
    var i, j int = 1, 2
    // 在函数中， := 简洁赋值语句在明确类型的地方，可以用于替代 var 定义
    // := 结构不能使用在函数外
    k := 3
    // 赋值之前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}
