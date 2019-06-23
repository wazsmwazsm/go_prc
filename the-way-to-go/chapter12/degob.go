package main

import (
	"fmt"
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	file, err := os.Open("vcard.gob")	
	if err != nil {
		log.Println("File open faild!")
	} 
	dec := gob.NewDecoder(file)
	var va VCard
	
	if err := dec.Decode(&va); err != nil {
		log.Println("Decode faild!")
	}
	fmt.Println(va.FirstName, va.LastName, va.Addresses[0].City, va.Addresses[1].City)
}