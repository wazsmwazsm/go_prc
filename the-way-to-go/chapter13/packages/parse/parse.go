package parse

import (
	"fmt"
	"strings"
	"strconv"
)

// A ParseError indicates an error in converting a word into an integer.
type ParseError struct {
	Index int // The index into the space-separated list of words.
	Word  string // The word that generated the parse error.
	Err   error  // The raw error that precipitated this error, if any.
}

func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

func Parse(input string) (numbers []int, err error) {
	defer func() {
		// 对外返回 err
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if ! ok {
				err = fmt.Errorf("pkg: %v", r) 
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	
	if len(fields) == 0 {
		panic("no words to parse") // 内部使用 panic
	}

	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err}) // 内部使用 panic
		}
		numbers = append(numbers, num)
	}

	return
}

