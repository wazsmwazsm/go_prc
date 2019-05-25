package main

import (
    "fmt"
    "math"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(100))
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
