package main

import "fmt"

// var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。
// var 语句可以定义在包或函数级别
var c, python, java bool

func main() {
    var i int
    fmt.Println(i, c, python, java)
}
