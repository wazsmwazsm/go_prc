package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "products.txt"
	outputFile := "products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
        panic(err.Error())
    }
}