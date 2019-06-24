package main

import (
	"fmt"
	"os"
	"log"
	"errors"
)

func main() {
	isErr := false
	content := "aaa"
	var errHandler func(content string)
	if ! isErr {
		errHandler = errorHandler(response)
	} else {
		errHandler = errorHandler(responseErr)
	}
	
	errHandler(content)
}

func response(content string) {
	if len(content) == 0 {
		panic(errors.New("Content is empty!"))
	}

	fmt.Fprintln(os.Stdout, content)
}

func responseErr(content string) {
	panic(errors.New(content))
}

func errorHandler(fn func(content string)) func(content string) {
	return func(content string) {
		defer func() {
			if err, ok := recover().(error); ok {
				log.Printf("run time panic: %v", err)
			}
		}()
		fn(content)
	}
}
