package main

import (
	"./packtest/greetings"
)

func main() {
	if greetings.IsAM() || greetings.IsAfternoon() {
		greetings.SayGoodDay()
	} else if greetings.IsEvening() {
		greetings.SayGoodNight()
	}
}