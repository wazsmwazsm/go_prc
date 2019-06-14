package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("Please enter your name:")
    input, err := inputReader.ReadString('\n')

    if err != nil {
        fmt.Println("There were errors reading, exiting program.")
        return
	}
	
	fmt.Printf("Your name is %s", input)

	switch input {
	case "Philip\n": fmt.Println("Welcome Philip!")
	case "Chris\n":   fmt.Println("Welcome Chris!")
    case "Ivo\n":     fmt.Println("Welcome Ivo!")
    default: fmt.Printf("You are not welcome here! Goodbye!")
	}
	// version 2:   
    switch input {
	case "Philip\n":  fallthrough
	case "Ivo\n":     fallthrough
	case "Chris\n":   fmt.Printf("Welcome %s\n", input)
	default: fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\n", "Ivo\n":   fmt.Printf("Welcome %s\n", input)
	default: fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}
