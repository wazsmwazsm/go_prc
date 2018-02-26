package main

import "fmt"

func main() {

    p := []byte{2, 3, 5}
    p = AppendByte(p, 7, 11, 13)

    fmt.Println(p)
}

// 将数据追加到切片的尾部
func AppendByte(slice []byte, data ...byte) []byte {
    m := len(slice)
    n := m + len(data)
    // 如果需要，扩充
    if n > cap(slice) {
        newSlice := make([]byte, (n + 1) * 2)
        copy(newSlice, slice)
        slice = newSlice
    }
    // 若果容量够，直接切片
    slice = slice[0:n]
    copy(slice[m:n], data)

    return slice
}
