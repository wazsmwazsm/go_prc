package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

type Test struct {
	Name    string
	Age     int
	Address []string
	Info    struct {
		Hobby  string
		Number int
	}
}

var yamlStr = `
name: Bob
age: 32
address:
- New York
- London
info:
  hobby: haha
  number: 4
`

func main() {
	// marshal
	t := Test{"Jack", 22, []string{"beijing chaoyang", "shanghai pudong"}, struct {
		Hobby  string
		Number int
	}{"Ping pong", 2}}

	res, err := yaml.Marshal(t)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s\n", res)

	// un marshal
	t2 := Test{}
	err = yaml.Unmarshal([]byte(yamlStr), &t2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(t2)

	// use map
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(yamlStr), &m)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(m)

	res, err = yaml.Marshal(m)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s\n", res)
}
