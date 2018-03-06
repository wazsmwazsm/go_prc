package main

import (
    "io"
    "os"
    "strings"
)

// 对数据流进行修改, 通过一个 io.Reader 包装另一个 io.Reader
type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {

    n, err = r.r.Read(p)

    for i := 0; i < n; i++ {
        p[i] = rot13(p[i])
    }
    
    return
}

func main()  {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

func rot13(b byte) byte {

    var a, z byte

    switch {
        case 'a' <= b && b <= 'z':
            a, z = 'a', 'z'
        case 'A' <= b && b <= 'Z':
            a, z = 'A', 'Z'
        default:
            return b
    }
    // 通过取余获取溢出后从头开始的位置
    return (b - a + 13) % (z - a + 1) + a
}
