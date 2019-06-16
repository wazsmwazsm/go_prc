package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
)

func main() {
	fd, err := os.Open("./test.txt")
	if err != nil {
        fmt.Printf("An error occurred on opening the inputfile\n" +
            "Does the file exist?\n" +
            "Have you got acces to it?\n")
        return // exit the function on error
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	for {
		inputString, readerError := reader.ReadString('\n')
		fmt.Println(inputString)
		if readerError == io.EOF {
			return
		}
	}
}