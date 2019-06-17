package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open(filepath.Base("products2.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 按照列读取
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil { // scans until empty line
            break
		}
		col1 = append(col1, v1)
        col2 = append(col2, v2)
        col3 = append(col3, v3)
	}

	fmt.Println(col1)
    fmt.Println(col2)
    fmt.Println(col3)
}
