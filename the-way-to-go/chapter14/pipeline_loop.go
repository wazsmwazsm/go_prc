// pipeline loop 版本, 性能好, 内存占用小, pipe 越多越有优势
package main

import (
	"errors"
	"fmt"
	"time"
)

// Handler A pipe Handler
type Handler func(interface{}) (interface{}, error)

// Pipeline A pipeline struct
type Pipeline struct {
	Handlers []Handler
}

// Add add a handler to pipeline
func (p *Pipeline) Add(handler Handler) {
	p.Handlers = append(p.Handlers, handler)
}

// Flow flow a pipeline
func (p *Pipeline) Flow(payload interface{}) (interface{}, error) {
	res := payload
	var err error
	for i := 0; i < len(p.Handlers); i++ {
		res, err = p.Handlers[i](res)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func main() {
	testPipe()
	testPipeBroke()
	testPipeMany(10000)
}

func testPipe() {
	p := new(Pipeline)
	p.Add(func(payload interface{}) (interface{}, error) {
		return payload.(int) + 2, nil
	})

	p.Add(func(payload interface{}) (interface{}, error) {
		return payload.(int) + 4, nil
	})

	if res, err := p.Flow(11); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.(int))
	}
}

func testPipeBroke() {
	p := new(Pipeline)
	p.Add(func(payload interface{}) (interface{}, error) {
		return payload.(int) + 2, nil
	})

	p.Add(func(payload interface{}) (interface{}, error) {
		return nil, errors.New("Broke pipe")
	})

	p.Add(func(payload interface{}) (interface{}, error) {
		return payload.(int) + 4, nil
	})

	if res, err := p.Flow(11); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.(int))
	}
}

func testPipeMany(count int) {
	start := time.Now()
	p := new(Pipeline)
	for i := 0; i < count; i++ {
		p.Add(func(payload interface{}) (interface{}, error) {
			return payload.(int) + i, nil
		})
	}

	if res, err := p.Flow(11); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.(int))
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}
