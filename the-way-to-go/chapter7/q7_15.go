package main

import "fmt"

func main() {
	s := []byte{'O', 'o', 'o', 'o', 'p', 's', '!', 'h', 'e', 'l', 'l', 'l', 'l', 'o', '!', '!'}

	s = unique(s)
	fmt.Println(string(s))
}

func unique(s []byte) (us []byte) {
	for i := 0; i < len(s); i++ {
		if i == 0 {
			us = append(us, s[i])
		} else if s[i] != s[i - 1] {
			us = append(us, s[i])
		}
	}

	return 
}

