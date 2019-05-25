package main

import "fmt"

// go 中都是传值传入，形参是一个拷贝
func changeSlice(s []int)  {
    s = nil
}
// slice 本身是传值，不过传入后引用的底层数组指针不变，所以这样会改变 slice 的值
func changeSliceElement(s []int)  {
    s[0] = 233
}
// 指针也是传值, 单纯改变指针没有效果
func changeSlicePointer(s *[]int)  {
    s = nil
}
// 指针是传值, 不过可以改变指针指向的值
func changeSlicePointerChangeValue(s *[]int)  {
    *s = []int{1, 1, 2}
}
func main() {
    s := []int{2, 3, 4, 5}
    fmt.Println(s) // [2 3 4 5]
    changeSlice(s)
    fmt.Println(s) // [2 3 4 5]

    // 简单的赋值引用的是同一个底层数组指针
    // 想要新的要使用 make copy
    s2 := s
    s3 := make([]int, len(s))
    copy(s3, s)
    fmt.Println(s) // [2 3 4 5]
    changeSliceElement(s)
    fmt.Println(s) // [233 3 4 5]
    fmt.Println(s2) // [233 3 4 5]
    fmt.Println(s3) // [2 3 4 5]

    fmt.Println(s) // [233 3 4 5]
    changeSlicePointer(&s)
    fmt.Println(s) // [233 3 4 5]

    fmt.Println(s) // [233 3 4 5]
    changeSlicePointerChangeValue(&s)
    fmt.Println(s) // [1, 1, 2]
}
