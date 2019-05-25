package main

import "fmt"

func main() {
    sum := 0
    // 注意，这里 i 只有 for 中可见
    for i := 0; i < 10; i++ {
        // var a = 10 // 同样只有 for 中可见 变量申明但不使用会报错
        sum += i
    }
    fmt.Println(sum)

    // 循环初始化语句和后置语句都是可选的
    sum = 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)

    // C 系语言的 while
    sum = 1
    for sum < 2000 {
        sum += sum
    }
    fmt.Println(sum)

    // 死循环
    for {      
    }
}
