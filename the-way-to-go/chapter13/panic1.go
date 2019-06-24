package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
}

func main() {
	fmt.Println(user)
}
