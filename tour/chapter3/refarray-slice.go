package main

import "fmt"

func main() {
    d := []byte{'r', 'o', 'a', 'd'}
    e := d[2:]
    fmt.Printf("%s\n", e)
    fmt.Printf("%s\n", d)
    // 切片会复用一个数组，不会进行拷贝
    e[1] = 'm'
    fmt.Printf("%s\n", e)
    fmt.Printf("%s\n", d)
}
