package main 

import "fmt"

type employee struct {
	salary int
}

func main() {
	jack := employee{5000}
	fmt.Println("涨薪前")
	fmt.Println(jack.salary)
	jack.giveRaise(30)
	fmt.Println("涨薪后")
	fmt.Println(jack.salary)
}

func (e *employee) giveRaise(percent int) {
	e.salary = e.salary + (e.salary * percent) / 100
}
