package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.baidu.com/",
	"http://example.com//",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
