package main

import (
	"fmt"
	"math"
	"errors"
)

func main() {
	fmt.Print("First example with -1: ")
	r1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", r1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", r1, err1)
	}

	fmt.Print("Than example with 5: ")
	r2, err2 := MySqrt(5)
	if err2 != nil {
		fmt.Println("Error! Return values are: ", r2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", r2, err2)
	}

	fmt.Println(MySqrt2(5))
}

func MySqrt(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("不接受负数哦~")
	}

	return math.Sqrt(f), nil
}

func MySqrt2(f float64) (r float64, err error) {
	if f < 0 {
		r, err = float64(math.NaN()), errors.New("不接受负数哦~")
	} else {
		r = math.Sqrt(f)
	}

	return
}
