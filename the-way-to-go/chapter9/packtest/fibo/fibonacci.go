package fibo

func Get(n int) int {
	x, y := 1, 0
	for i := 1; i <= n; i++ {
		x, y = y, x + y 
	}

	return y
}

func GetByOperator(n int, operator string) int {
	x, y := 1, 0
	for i := 1; i <= n; i++ {
		switch operator {
		case "+":
			x, y = y, x + y 
		case "*":
			x, y = y, x * y 
		}
	}

	return y
}