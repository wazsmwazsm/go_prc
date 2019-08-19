package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8)
	after := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-after:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
