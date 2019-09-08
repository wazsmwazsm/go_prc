package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the url...")
	input, err := inputReader.ReadString('\n')
	input = strings.Trim(input, "\r\n")
	resp, err := http.Get(input)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}
