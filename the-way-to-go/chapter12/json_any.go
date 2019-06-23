package main

import (
	"fmt"
	"encoding/json"
	"os"
)

func main() {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)

	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Json decode faild! %s\n", err)
		os.Exit(1)
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}	
		default:
			fmt.Println(k, "is of a type I don’t know how to handle")
		}
	}
}