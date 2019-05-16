package main

import "fmt"

func main() {
    s := "\u00ff\u754c"
    for i, c := range s {
        fmt.Printf("%d:%c ", i, c) // 0:ÿ 2:界 这里 Unicode 占用多字节
	}
	fmt.Printf("%c %c %c\n", s[0], s[1], s[2]) // 获取的只是字节 (ascii 字符串时会获取到字符)
	fmt.Printf("%s %s", s[:2], s[2:5]) // 0:ÿ 2:界
}
