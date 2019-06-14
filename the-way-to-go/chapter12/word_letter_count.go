package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter something:")
	input, _ := inputReader.ReadString('S')
	length := len(input) - strings.Count(input, "\r") - strings.Count(input, "\n")
	wordCount := len(strings.Fields(input))
	lines := len(strings.Split(input, "\n"))
	fmt.Println(length, wordCount, lines)
}
