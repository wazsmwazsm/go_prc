package main

import (
	"fmt"
	"os"
)

func main() {
	who := "Hello "
	if len(os.Args) > 1 {
		who += os.Args[1]
	}
	fmt.Println(who)
}
