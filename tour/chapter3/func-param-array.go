package main

import "fmt"

// array 是值传递
func changeArray(a [2]int)  {
    t := [2]int{2, 5}
    a = t
}

// 和 slice 、map 不同, 传入的数组整体是一个值
// 修改数组的元素这里是不会奏效的
func changeArrayElement(a [2]int)  {
    a[1] = 5
}
// 数组的指针的话，就是普通指针传值的效果
// 修改指针本身没什么用
func changeArrayPointer(a *[2]int)  {
    t := [2]int{2, 5}
    a = &t
}
// 修改数组指针指向是肯定奏效的
// 这点对任何指针都有效
func changeArrayPointerChangeValue(a *[2]int)  {
    *a = [2]int {2, 5}
}

func main() {
    a := [2]int {1, 2}
    fmt.Println(a)
    changeArray(a)
    fmt.Println(a)

    a1 := a

    fmt.Println(a)
    changeArrayElement(a)
    fmt.Println(a)
    fmt.Println(a1)

    fmt.Println(a)
    changeArrayPointer(&a)
    fmt.Println(a)

    a2 := a
    fmt.Println(a)
    changeArrayPointerChangeValue(&a)
    fmt.Println(a)
    fmt.Println(a2)
}
