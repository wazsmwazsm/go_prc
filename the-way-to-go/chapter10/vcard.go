package main 
import "fmt"

type Address struct {
	title string
}
type VCard struct {
	name string
	birthDay string
	addr Address
}

func main() {
	c1 := VCard{"Bob", "1992-02-01", Address{"Beijing chaoyang 1122"}}
	fmt.Printf("name: %s, brith day: %s, address: %s", c1.name, c1.birthDay, c1.addr.title)
}
