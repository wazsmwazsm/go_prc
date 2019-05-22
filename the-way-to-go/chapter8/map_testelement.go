package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	
	if v, ok := map1["Beijing"]; ok {
		fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", v)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}

	v, ok := map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ?: %t\n", ok)
	fmt.Printf("Value is: %d\n", v)

	// delete an item
	delete(map1, "Washington")
	
	if v, ok := map1["Washington"]; ok {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", v)
	} else {
		fmt.Printf("map1 does not contain Washington")
	}
}
