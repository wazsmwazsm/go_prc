package main

import (
	"encoding/gob"
	"os"
	"log"
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
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	va := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	// 使用编码流:
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE | os.O_WRONLY, 0666)
	defer file.Close()

	enc := gob.NewEncoder(file)
	err := enc.Encode(va)
	if err != nil {
		log.Println("Error in encoding gob")
	}


}