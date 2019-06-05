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

func main() {
	var s Simpler = &Person{"Jack"}
	
	fmt.Println(s.Get())
	s.Set("Mark")
	fmt.Println(s.Get())
}
