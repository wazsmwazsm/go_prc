package main

import (
	"github.com/wazsmwazsm/uc"
	"greetings"
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
