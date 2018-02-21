
package main

import "fmt"

func fibonacci() func() int {
    // 保存数列
    var s []int
    // 计次
    i := 0

    return func() int {
        // 当前的索引
        current := i
        i++
        // 判断是否是数列开头
        if (current == 0 || current == 1) {
            s = append(s, current)
        } else {
            s = append(s, s[current-2] + s[current-1])
        }
        // 返回当前索引下数列的值
        return s[current]
    }
}

func main()  {
    f := fibonacci()

    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
