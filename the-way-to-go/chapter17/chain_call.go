package main

import (
	"fmt"
)

type Num interface {
	Add(int) Num
	Sub(int) Num
}

type number int

func (n number) Add(num int) Num {
	return n + number(num)
}

func (n number) Sub(num int) Num {
	return n - number(num)
}

func main() {
	var n number = 10
	var i Num
	i = n

	fmt.Println(i.Add(22).Sub(2).Add(5))
}
