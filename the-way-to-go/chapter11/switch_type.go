package main

import "fmt"

func main() {
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
}

// 参数是空接口类型的变长参数
func classifier(items ...interface{}) {
	for i, x := range items {
		switch v := x.(type) {
		case bool:
			fmt.Printf("Param #%d - %t is a bool\n", i, v)
		case float64:
			fmt.Printf("Param #%d - %f is a float64\n", i, v)
		case int, int64:
			fmt.Printf("Param #%d - %v is a int\n", i, v)
		case nil:
			fmt.Printf("Param #%d - %s is a nil\n", i, v)
		case string:
			fmt.Printf("Param #%d - %s is a string\n", i, v)
		default:
			fmt.Printf("Param #%d - %v is unknown\n", i, v)
		}
	}
}
