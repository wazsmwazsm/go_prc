package main

import (
    "golang.org/x/tour/pic"
    "math"
)

func Pic(dx, dy int) [][]uint8 {
    var s [][]uint8

    for i := 0; i < dy; i++ {
        var tmp []uint8
        for j := 0; j < dx; j++ {
            pixel := float64(i) * math.Log(float64(j))
            tmp = append(tmp, uint8(pixel))
        }

        s = append(s, tmp)
    }

    return s
}

func main() {
    pic.Show(Pic)
}
