package main

import "fmt"

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func main() {
	for h := 0; h < HEIGHT; h++ {
		for w := 0; w < WIDTH; w++ {
			screen[w][h] = pixel((h + w) * 2)
		}
	}

	showScreen(&screen)
}

func showScreen(s *[WIDTH][HEIGHT]pixel) {
	var tmp string
	for h := 0; h < HEIGHT; h++ {
		tmp = ""
		for w := 0; w < WIDTH; w++ {
			tmp = fmt.Sprintf("%s %v ", tmp, s[w][h])
		}
		fmt.Println(tmp)
	}
}
