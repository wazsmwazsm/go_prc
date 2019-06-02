package main

import (
	"fmt"
	"./packages/structPack"
)

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16. // 省略写法
	// struct1.si1 = 6  // 小写成员对外部隐藏

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)
	// fmt.Printf("Mf1 = %f\n", struct1.si1)
	
}
