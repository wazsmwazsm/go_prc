package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

// 声明 map 键为 string 类型, 值为 Vertex 类型
var m map[string]Vertex

func main() {
    // 生成一个非空 map
    m = make(map[string]Vertex)
    // 赋值
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }

    fmt.Println(m["Bell Labs"])
}
