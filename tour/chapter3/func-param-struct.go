package main

import "fmt"

type Vertex struct {
    s []string
}

// 结构体也是值传递
func changeStruct(v Vertex)  {
    v = struct {s []string}{
        []string{"1", "b", "3"},
    }
}

// 结构体操作不会生效, 但是传入指针可以修改
func changeStructElement(v Vertex)  {
    v.s = []string{"a", "hello", "qin"}
}

// 这样可行
func changeStructPointerElement(v *Vertex)  {
    v.s = []string{"a", "hello", "qin"}
}

// 操作结构体的指针和普通指针是一样的
func changeStructPointer(v *Vertex)  {
    v = &Vertex{
        []string{"1", "b", "3"},
    }
}
// 修改结构体指针的指向，使它指向另一个结构体数据
func changeStructPointerChangeValue(v *Vertex)  {
    *v = Vertex{
        []string{"1", "b", "3"},
    }
}
func main() {
    v := Vertex{
        []string{"a", "b", "c"},
    }
    fmt.Println(v)
    changeStruct(v)
    fmt.Println(v)

    fmt.Println(v)
    changeStructPointerElement(&v)
    fmt.Println(v)
    // 结构体的赋值是拷贝, 不会同时引用一份数据
    v1 := v
    fmt.Println(v)
    changeStructElement(v)
    fmt.Println(v)
    fmt.Println(v1)

    fmt.Println(v)
    changeStructPointer(&v)
    fmt.Println(v)

    fmt.Println(v)
    changeStructPointerChangeValue(&v)
    fmt.Println(v)
}
