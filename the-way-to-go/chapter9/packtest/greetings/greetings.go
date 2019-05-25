package greetings

import (
	"fmt"
	"time"
)

func SayGoodDay() {
	fmt.Println("Good Day!")
}

func SayGoodNight() {
	fmt.Println("Good Night!")
}

func IsAM() bool {
	t := time.Now()

	if t.Hour() <= 12 {
		return true
	} else {
		return false
	}
}

func IsAfternoon() bool {
	t := time.Now()

	if t.Hour() > 12 && t.Hour() < 19 {
		return true
	} else {
		return false
	}
}

func IsEvening() bool {
	t := time.Now()

	if t.Hour() > 19 {
		return true
	} else {
		return false
	}
}