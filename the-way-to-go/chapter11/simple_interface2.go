package main

import "fmt"

type Simpler interface {
	Get() string
	Set(string)
}

type Person struct {
	name string
}

func (p *Person) Get() string {
	return p.name
}

func (p *Person) Set(name string) {
	p.name = name
}

func gI(any interface{}) {
	switch v := any.(type) {
	case *Person:
		fmt.Println(v.Get())
	default:
		fmt.Println("The type is not safe!")
	}
}

func main() {
	s := &Person{"Jack"}
	
	gI(s)
	i := 1
	gI(i)
}
