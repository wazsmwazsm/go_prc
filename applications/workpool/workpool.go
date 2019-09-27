package workpool

import (
	"sync/atomic"
)

type Task func()

type Pool struct {
	capacity int64
	running  int64
	taskC    chan Task
}

func NewPool(capacity int64) *Pool {
	return &Pool{
		capacity: capacity,
		taskC:    make(chan Task),
	}
}

func (p *Pool) Run(task Task) {
	if p.running < p.capacity {
		go func() {
			for {
				select {
				case t := <-p.taskC:
					t()
				}
			}
		}()
		atomic.AddInt64(&p.running, 1)
	}
	p.taskC <- task
}
