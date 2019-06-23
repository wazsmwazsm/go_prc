package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
)

func main() {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		outputString := string(inputString[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		outputWriter.Flush() // 源程序没有 flush buf 到文件中, 此行修复
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Conversion done")
}