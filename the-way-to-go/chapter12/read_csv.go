package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"path/filepath"
)

func main() {
	file, err := os.Open(filepath.Base("products3.txt"))
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	fmt.Println("Title Price Quantity\n")
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")
		if err == io.EOF {
			break
		}
		lineSlice := strings.Split(line, ";")
		
		for _, l := range lineSlice {
			fmt.Printf("%s", l)
		}
		fmt.Println()
	}
}
