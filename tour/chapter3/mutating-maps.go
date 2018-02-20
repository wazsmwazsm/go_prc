package main

import "fmt"

func main() {
  	m := make(map[string]int)

    // 修改 map
    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])
    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])
    // 删除元素
    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])

    // 获取值，检查键是否存在
    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}
