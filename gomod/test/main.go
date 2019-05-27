package main

import (
	"github.com/wazsmwazsm/uc"
	"github.com/wazsmwazsm/greetings"
	// "./greetings" // 这样是不行的, 本地宝需要 go.mod 里写 replace
	"fmt"
)

func main() {
	str1 := "USING package uc!"
	fmt.Println(uc.UpperCase(str1))

	if greetings.IsAM() || greetings.IsAfternoon() {
		greetings.SayGoodDay()
	} else if greetings.IsEvening() {
		greetings.SayGoodNight()
	}
}
