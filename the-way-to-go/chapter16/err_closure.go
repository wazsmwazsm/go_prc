package main

import (
	"errors"
	"fmt"
)

func main() {
	test(1, 2)
	test(6, 2)
	test(1, 7)
}

func test(a, b int) {
	// 闭包集中处理错误
	err := func() error {
		// 这里只是演示，其实这种情况直接 a > 5 || b > 5 最合适
		if a > 5 {
			return errors.New("a must not grater then 5")
		}

		if b > 5 {
			return errors.New("b must not grater then 5")
		}

		return nil
	}()

	if err != nil {
		fmt.Println(err)
	}
}
