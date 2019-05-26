// greetings. say good day\night
package greetings

import (
	"fmt"
	"time"
)

// SayGoodDay, 当前是白天，打印 Good Day
func SayGoodDay() {
	fmt.Println("Good Day!")
}

// SayGoodNight, 当前是白天，打印 Good Night
func SayGoodNight() {
	fmt.Println("Good Night!")
}

// IsAM 判断是不是上午
func IsAM() bool {
	t := time.Now()

	if t.Hour() <= 12 {
		return true
	} else {
		return false
	}
}
// IsAfternoon 判断是不是下午
func IsAfternoon() bool {
	t := time.Now()

	if t.Hour() > 12 && t.Hour() < 19 {
		return true
	} else {
		return false
	}
}
// IsEvening 判断是不是晚上
func IsEvening() bool {
	t := time.Now()

	if t.Hour() > 19 {
		return true
	} else {
		return false
	}
}