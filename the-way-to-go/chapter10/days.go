package main

import (
	"fmt"
)

var week [7]string = [7]string{
	"Monday", 
	"Tuesday", 
	"Wednesday", 
	"Thursday", 
	"Friday", 
	"Saturday",
	"Sunday",
}

type Day int

func (d Day) String() string {
	return week[d - 1]
}

func main() {
	var i Day = 1
	for ; i <= 7; i++ {
		fmt.Printf("%v\n", i)
	}
}
