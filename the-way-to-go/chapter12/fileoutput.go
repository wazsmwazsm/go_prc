package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string

	// os.O_RDONLY：只读
	// os.O_WRONLY：只写
	// os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
	// os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为0。
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY | os.O_CREATE, 0644)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return  
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile) // 定义缓冲区（使用缓冲区写文件可以减少 io 操作）
	outputString := "hello world!\n"
	for i:=0; i<10; i++ {
		outputWriter.WriteString(outputString) // 字符串写入缓冲区
	}
	outputWriter.Flush() // 把缓冲区的数据写入到文件（1 次 io 操作）
}
