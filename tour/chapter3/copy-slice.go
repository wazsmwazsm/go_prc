package main

import "fmt"

func main() {
    s := make([]byte, 5)
    fmt.Println("\ncap before: ", cap(s))
    // 通过新建复制的方式扩充容量
    t := make([]byte, len(s), (cap(s) + 1) * 2)
    // copy
    // for i := range s {
    //     t[i] = s[i]
    // }
    copy(t, s)
    s = t

    fmt.Println("\ncap after: ", cap(s))
}
