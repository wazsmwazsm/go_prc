package main
import "fmt"

type number struct {
	f float32
}

type nr number

func main() {
	a := number{5.0}
	b := nr{5.0}
	// var i float32 = b   // compile-error: cannot use b (type nr) as type float32 in assignment
    // var i = float32(b)  // compile-error: cannot convert b (type nr) to type float32
	// var c number = b    // compile-error: cannot use b (type nr) as type number in assignment
	var c = number(b)
	fmt.Println(a, b, c)
}
