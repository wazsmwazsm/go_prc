package main

import "fmt"

func main() {
	a := [4]float64{7.0, 9.33, 9.2, 3.13}

	fmt.Printf("The sum of the array is: %f", sum(&a))
}

// 使用命名返回值时，sum 相当于已经声明，在函数体中不用再声明
func sum(a *[4]float64) (sum float64) {
	for _, v := range a {
		sum += v
	} 

	return
}
