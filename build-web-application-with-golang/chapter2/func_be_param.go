package main

import "fmt"
// 声明一个函数类型
type testInt func(int) bool

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer % 2 == 0 {
		return true
	}
	return false
}

// 函数做参数
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main(){
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)    // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)  // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}
