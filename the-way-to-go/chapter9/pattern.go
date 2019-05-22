package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pattern := "[0-9]+.[0-9]+"

	// also regexp.MatchString(pattern, searchIn)
	if ok, _ := regexp.Match(pattern, []byte(searchIn)); ok {
		fmt.Println("Matched!")
	}

	// 通过 Compile 方法返回一个 Regexp 对象, 进行更灵活的操作
	re, _ := regexp.Compile(pattern)
	// 替换全部 
	str := re.ReplaceAllString(searchIn, "**.**") 
	fmt.Println(str)

	// 通过函数替换
	str2 := re.ReplaceAllStringFunc(searchIn, func (s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v * 2, 'f', 3, 32)
	})

	fmt.Println(str2)
}
