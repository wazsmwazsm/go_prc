package workpool

import (
	"context"
	"sync/atomic"
)

type Task func()

type Pool struct {
	capacity uint64
	running  uint64
	taskC    chan Task
	ctx      context.Context
	quit     context.CancelFunc
}

func NewPool(capacity uint64) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	return &Pool{
		capacity: capacity,
		taskC:    make(chan Task),
		ctx:      ctx,
		quit:     cancel,
	}
}

func (p *Pool) Fill() {
	if p.running < p.capacity {
		atomic.StoreUint64(&p.running, p.capacity)
		for i := 0; i < int(p.capacity); i++ {
			go func(ctx context.Context) {
				for {
					select {
					case t := <-p.taskC:
						t()
					case <-ctx.Done():
						return
					}
				}
			}(p.ctx)
		}
	}

}

func (p *Pool) Run(task Task) {
	if p.running < p.capacity {
		go func(ctx context.Context) {
			for {
				select {
				case t := <-p.taskC:
					t()
				case <-ctx.Done():
					return
				}
			}
		}(p.ctx)
		atomic.AddUint64(&p.running, 1)
	}
	p.taskC <- task
}

func (p *Pool) Release() {
	p.quit()
}
