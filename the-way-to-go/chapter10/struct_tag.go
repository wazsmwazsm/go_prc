package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	filed1 bool "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	t := TagType{true, "Barak Obama", 2}
	for i := 0; i < 3; i++ {
		refTag(t, i)
	}
}

func refTag(t TagType, ix int) {
	tType := reflect.TypeOf(t)
	ixField := tType.Field(ix)
	fmt.Printf("Name:%v, Tag:%v, Package:%v, Type:%v\n", ixField.Name, ixField.Tag, ixField.PkgPath, ixField.Type)
}
