package main

import (
	"fmt"
	"bufio"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin) // 绑定 inputReader 到 stdin
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n') // 遇到 \n 终止读取
	if err == nil {
        fmt.Printf("The input was: %s\n", input)
    }
}

