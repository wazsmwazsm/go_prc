package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Person 实现了 Stringer 接口, 重写了 String 方法
func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main()  {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}
	  fmt.Println(a, z)
}
