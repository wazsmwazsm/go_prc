package main

import (
    "fmt"
    "io"
    "strings"
)

func main()  {
    r := strings.NewReader("Hello, Reader!")
    // 创建一个切片作为 Read 函数的读取 buffer
    b := make([]byte, 8) // 每次 8 字节的速度读取它的输出

    for {
        // Read 方法返回填充的字节数和错误值
        n, err := r.Read(b)
        // 这里打印 b 的值发现 buffer 每次有新值则从左往右覆盖
        // 每次读取不会擦除
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n]) // 读取 buffer 的方式
        // 遇到数据流的结尾时, 会返回一个 io.EOF 错误
        if err == io.EOF {
            break
        }
    }

}
