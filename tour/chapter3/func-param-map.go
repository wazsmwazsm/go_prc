package main

import "fmt"

// map 也是值传递
func changeMap(m map[int]string)  {
    m = nil
}

// map 变量更像是指针，传入的参数可以直接修改 map 的值
// 因为其底层引用的数据是同一份
func changeMapElement(m map[int]string)  {
    m[5] = "shut up"
}
// 操作 map 的指针和普通指针是一样的
func changeMapPointer(m *map[int]string)  {
    m = nil
}
// 修改 map 指针的指向，使它指向另一个 map 数据
func changeMapPointerChangeValue(m *map[int]string)  {
    *m = map[int]string {
      1: "a",
      2: "b",
      5: "c",
    }
}
func main() {
    m := map[int]string {
      1: "hello",
      2: "hi",
      5: "come on",
    }
    fmt.Println(m)
    changeMap(m)
    fmt.Println(m)

    m1 := m

    fmt.Println(m)
    changeMapElement(m)
    fmt.Println(m)
    fmt.Println(m1)

    fmt.Println(m)
    changeMapPointer(&m)
    fmt.Println(m)

    fmt.Println(m)
    changeMapPointerChangeValue(&m)
    fmt.Println(m)
}
