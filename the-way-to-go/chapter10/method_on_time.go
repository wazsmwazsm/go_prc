package main

import (
	"fmt"
	"time"
)

// 使用匿名字段嵌入，这样可以使接受者类型和方法在一个包中
type myTime struct {
	time.Time 
}

func main() {
	m := myTime{time.Now()}
	// 调用匿名Time上的String方法
	fmt.Println("Full time now:", m.String())
	// 调用myTime.first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}
