package main

import "fmt"

type obj interface{}

type objs []obj

func main() {
	mapFunc("hai", 11, "hi", "baby", 4)
}

func mapFunc(args ...interface{}) {
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			fmt.Println(v * 2)
		case string:
			fmt.Println(v + v)
		}
	}
	
}
