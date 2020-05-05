package main

// 迭代的方式
import "fmt"

func main() {
	fmt.Println(fibonacci(10))
}

// 其实也是 DP 思想，只不过进一步优化省掉了额外的 DP 数组
func fibonacci(n int) (res []int) {
	a, b := 1, 0 // 斐波那契的值 f(n) = f(n-1) + f(n-2), 那么维护 f(n-1) + f(n-2) 两个值即可
	for i := n; i > 0; i-- {
		// b 作为 f(n-1) a 为 f(n-2)，a+b 就是 f(n)，每轮计算后 b 的值就是 f(n)
		// 然后 b 又作为下轮计算的 f(n-1)，一次迭代
		a, b = b, a+b
		res = append(res, b)
	}

	return res
}
