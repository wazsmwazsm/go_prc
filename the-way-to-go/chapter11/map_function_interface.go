package main

import "fmt"

type obj interface{}

type objs []obj

func main() {
	obj1 := objs([]obj{"hai", 11, "hi", "baby", 4})
	mapFunc(obj1)
	fmt.Println(obj1)
}

func mapFunc(o objs) {
	for i, obj := range o {
		switch v := obj.(type) {
		case int:
			o[i] = v * 2
		case string:
			o[i] = v + v
		}
	}
}
