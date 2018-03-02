package main

import (
    "golang.org/x/tour/pic"
    "math"
)

// func Pic(dx, dy int) [][]uint8 {
//     var s [][]uint8
//
//     for i := 0; i < dy; i++ {
//         var tmp []uint8
//         for j := 0; j < dx; j++ {
//             pixel := float64(i) * math.Log(float64(j))
//             tmp = append(tmp, uint8(pixel))
//         }
//
//         s = append(s, tmp)
//     }
//
//     return s
// }

func Pic(dx, dy int) [][]uint8 {
    // 构造 slice
    p := make([][]uint8, dy)

    for i := range p {
        p[i] = make([]uint8, dx)
    }
    // 填充数据
    for y, row := range p {
        for x := range row {
            row[x] = uint8(x * y)
        }
    }

    return p
}


func main() {
    pic.Show(Pic)
}
