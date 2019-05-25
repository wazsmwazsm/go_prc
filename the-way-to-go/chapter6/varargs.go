package main

import "fmt"

func main() {
	typecheck(12, "good luck!", 4.332, true, []int{1,2})
}

// 使用默认的空接口 interface{}，这样就可以接受任何类型的参数
func typecheck(values ...interface{}) {
	for _, value := range values {
		// 使用 switch 判断类型， value 必须是接口值
		switch v := value.(type) {
		case int:
			fmt.Printf("integer: %d\n", v)
		case float64:
			fmt.Printf("float64: %f\n", v)
		case string:
			fmt.Printf("string: %s\n", v)
		case bool:
			fmt.Printf("boolean: %t\n", v)
		default:
			fmt.Printf("othertype: %v\n", v)
		}
	}
}
