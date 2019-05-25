package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4

    fmt.Println(v.X)

    // 结构体的赋值是拷贝, 不会指向同一个值
    a := v
    a.X = 5
    fmt.Println(a)
    fmt.Println(v)
}
