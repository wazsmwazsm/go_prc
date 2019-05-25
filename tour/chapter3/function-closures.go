package main

import "fmt"

// 返回一个函数作为闭包
func adder() func(int) int {
    sum := 0
    // 每个闭包绑定一个 sum 变量
    // 实现类似其他语言 static 属性
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()

    for i := 0; i < 10; i++ {
        fmt.Println(pos(i), neg(-2*i))
    }

}
