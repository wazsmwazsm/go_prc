package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"flag"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 { // 没有传参，直接绑定 stdin
		cat(bufio.NewReader(os.Stdin)) 
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i)) // 取参数做文件名
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}

}
