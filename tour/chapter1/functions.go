package main

import "fmt"

func add(x int, y int) int {
    return x + y
}
// 当两个或多个连续的函数命名参数是同一类型，
// 则除了最后一个类型之外，其他都可以省略
func add2(x, y int) int {
    return x + y
}

// 函数可以返回任意数量的返回值。
func swap(x, y string) (string, string) {
    return y, x
}

// Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。
// 返回值的名称应当具有一定的意义，可以作为文档使用。
// 没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。
// 直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x

    // 裸 返回
    return
    // return y, x 亦可以带参数，只要满足返回值声明
}

func main() {
    fmt.Println(add(42, 13))
    fmt.Println(add2(42, 55))
    // 相当于直接声明且初始化 (要求变量的类型确定)
    a, b := swap("hello", "world")
    fmt.Println(a, b)
    fmt.Println(split(17))
}
