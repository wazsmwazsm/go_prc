package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// unbufferd
	fmt.Fprintf(os.Stdout, "%s\n", "Hello world! - unbufferd")
	// buffered: os.Stdout implements io.Writer
	buf := bufio.NewWriter(os.Stdout)
	// and now so does buf.
	fmt.Fprintf(buf, "%s\n", "Hello world! - bufferd")
	buf.Flush() // 从 buf 写入到 io
}