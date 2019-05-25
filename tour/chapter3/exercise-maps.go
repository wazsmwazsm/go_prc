package main

import (
  	"golang.org/x/tour/wc"
  	"strings"
)

func WordCount(s string) map[string]int {

  	m := make(map[string]int)

  	slice := strings.Fields(s)

  	// for i := 0; i < len(slice); i++ {
  	// 	m[slice[i]] += 1
  	// }

    for _, v := range slice {
        m[v] += 1
    }

  	return m
}

func main() {
    wc.Test(WordCount)
}
