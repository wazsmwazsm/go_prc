package main

import "fmt"

type Base struct {
	id int
}

func (this Base) Id() int {
	return this.id
}

type Person struct {
	*Base
	FirstName, LastName string
}

type Employee struct {
	*Person
	salary float64
}


func main() {
	e := &Employee{&Person{&Base{12}, "Jack", "Jons"}, 2500.00}

	fmt.Println(e.Id())
}
