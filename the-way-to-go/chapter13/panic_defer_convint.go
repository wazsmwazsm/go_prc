package main

import (
	"fmt"
	"math"
)

func main() {
	l := int64(15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err.Error())
	} else {
		fmt.Printf("[%T]:%d converted to an int32 is [%T]:%d", l, l, i, i)
	}
	fmt.Println()
	l = int64(math.MaxInt32 + 15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err.Error())
	} else {
		fmt.Printf("[%T]:%d converted to an int32 is [%T]:%d", l, l, i, i)
	}
}

func IntFromInt64(num int64) (result int32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panicking: ", r)
		}
	}()

	return ConvertInt64ToInt(num), nil
}

func ConvertInt64ToInt(num int64) (result int32) {
	if math.MinInt32 <= num && num <= math.MaxInt32 {
		return int32(num)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", num))
}
