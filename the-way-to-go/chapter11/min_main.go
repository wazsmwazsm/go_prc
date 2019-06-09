package main

import (
	"fmt"
	"./packages/min"
)

type person struct {
	name string
	age int
}

type persons []person

func (p persons) Len() int {
	return len(p)
}

func (p persons) Less(i, j min.Element) bool {

	if i.(person).age < j.(person).age {
		return true
	}
	
	return false
}

func (p persons) Get(index int) min.Element {
	var el min.Element = p[index]

	return el
}

func MinPersons(p persons) person {

	return min.Min(p).(person)
}

func main() {
	testInts()
	testFloats()
	testPersons()
}

func testInts() {
	ints := []int{8, 52, 4, 2, 12, 5, 7}
	min_ints := min.MinInts(ints)
	fmt.Printf("The min of the ints %v is %v\n", ints, min_ints)

}

func testFloats() {

	floats := []float64{1.2, 2.1, 0.23, 7.52}
	min_floats := min.MinFloats(floats)
	fmt.Printf("The min of the floats %v is %v\n", floats, min_floats)
}

func testPersons() {
	jack := person{"jack", 22}
	bob := person{"bob", 18}
	robber := person{"robber", 23}
	smith := person{"smith", 30}

	pers := persons{jack, bob, robber, smith}
	min_persons := MinPersons(pers)
	fmt.Printf("The min of the persons %v is %v\n", pers, min_persons)

}