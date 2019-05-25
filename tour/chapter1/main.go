package main

import (
    "fmt"
    "os"
)

func main() {
    for i, v := range os.Args {
        fmt.Printf("os.Args[%d] = %s\n", i, v)
    }
}
