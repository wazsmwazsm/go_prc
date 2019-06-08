package main

import (
	"fmt"
	"./packages/sort"
)

type Person struct {
	firstName, lastName string
}

type Persons []*Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return p[i].firstName < p[j].firstName
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	sortPersons()
}

func sortPersons() {
	Bob := Person{"Bob", "Clark"}
	Smith := Person{"V", "Smith"}
	Jack := Person{"Sex", "Jack"}

	p := Persons{&Jack, &Bob, &Smith}

	sort.Sort(p)
	if ! sort.IsSorted(p) {
		panic("sort fails!")
	}

	for _, v := range p {
		fmt.Printf("%s %s -- ", v.firstName, v.lastName)
	}
	fmt.Printf("\n")
}
