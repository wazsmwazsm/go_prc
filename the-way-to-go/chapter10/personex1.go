package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p Person) {
	// 值传递结构体的话，无法改变结构体的成员
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}


func main() {
	// 1-struct as a value type:
	var per1 Person
	per1.firstName = "Chris"
	per1.lastName = "Woodward"
	upPerson(per1)
	fmt.Printf("The name of the person is %s %s\n", per1.firstName, per1.lastName)

	// 2—struct as a pointer:
	per2 := new(Person)
	per2.firstName = "Chris"
	per2.lastName = "Woodward"
	(*per2).lastName = "Sam"  // 值的方式访问属性
	upPerson(*per2)
	fmt.Printf("The name of the person is %s %s\n", per2.firstName, per2.lastName)

	// 3—struct as a literal:
	per3 := &Person{"Chris", "Pop"}
	upPerson(*per3)
	fmt.Printf("The name of the person is %s %s\n", per3.firstName, per3.lastName)
}
