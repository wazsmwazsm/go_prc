package main

import "fmt"

func main() {
	s := "Hello My Friend!"
	s = reverseByte(s)
	fmt.Println(s)

	s = "你好朋友！"
	s = reverseRune(s)
	fmt.Println(s)
}

func reverseByte(s string) string {
	sb := []byte(s)
	var tmp byte
	for i := 0; i < len(sb) / 2; i++ {
		tmp = sb[i]
		sb[i] = sb[len(sb) - 1 - i]
		sb[len(sb) - 1 - i] = tmp
	}

	return string(sb)
}

func reverseRune(s string) string {
	sr := []rune(s)
	var tmp rune
	for i := 0; i < len(sr) / 2; i++ {
		tmp = sr[i]
		sr[i] = sr[len(sr) - 1 - i]
		sr[len(sr) - 1 - i] = tmp
	}

	return string(sr)
}
