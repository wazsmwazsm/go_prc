// pipeline loop 版本, 性能一般只做演示, 内存占用高, pipe 越多越慢
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

	// create goroutine chain
	firstIn := make(chan interface{})
	quit := make(chan error)
	var in, out chan interface{} = nil, firstIn
	for i := 0; i < len(p.Handlers); i++ {
		in, out = out, make(chan interface{})
		go pipe(in, out, quit, p.Handlers[i])
	}
	// start flow
	firstIn <- payload
	for {
		select {
		case res := <-out: // flow done
			return res, nil
		case err := <-quit: // flow broke
			return nil, err
		}
	}
}

// pipe, get data from in channel, handle data send to out channel
// if handle error, broke pipe
func pipe(in, out chan interface{}, quit chan error, handler Handler) {
	res, err := handler(<-in)
	if err != nil {
		quit <- err
		return
	}
	out <- res
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
