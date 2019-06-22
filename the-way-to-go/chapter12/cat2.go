package main 

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(os.Stdin)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not open %s, err: %s", flag.Arg(i), err)
			os.Exit(1)
		}
		cat(f)
		f.Close()
	}
}

func cat(f *os.File) {
	const MBUF = 512
	buf := make([]byte, MBUF) // 创建一个 buffer
	for {
		// 根据读取的字符数判断是否结束（遇到空行）
		switch nr, err := f.Read(buf); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
		case nr == 0:
			return
		case nr > 0:
			if nw, ew :=os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}