package main

import "fmt"

type Engine struct {
	engineName string
}

type Car struct {
	Engine
	wheelCount int
}

func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

type Mercedes struct {
	Car
	name string
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Printf("I am %s, with %v wheels, %s engine", m.name, m.numberOfWheels(), m.engineName)
}
func main() {
	m := Mercedes{Car{Engine{"V8"}, 4}, "Benz"}

	m.sayHiToMerkel()
}
