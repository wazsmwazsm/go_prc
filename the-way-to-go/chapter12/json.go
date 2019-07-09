package main

import (
	"fmt"
	"encoding/json"
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
	// JSON format:
	js, _ := json.Marshal(va)
	fmt.Printf("JSON format: %s\n", js)
	// JSON unformat:
	var va1 VCard
	json.Unmarshal(js, &va1)
	fmt.Printf("JSON parse: %s\n", va1)

	// 使用编码流:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE | os.O_WRONLY, 0666)
	defer file.Close()

	enc := json.NewEncoder(file)
	err := enc.Encode(va)
	if err != nil {
		log.Println("Error in encoding json")
	}

}
