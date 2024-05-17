package pool

import (
	"sync"
)

type Putter interface {
	Put()
}

func Put(data any) {
	putter, ok := data.(Putter)
	if ok {
		putter.Put()
	}
}

type Pool[T any] struct {
	pool sync.Pool
}

func New[T any](fn func() T) Pool[T] {
	return Pool[T]{
		pool: sync.Pool{New: func() interface{} { return fn() }},
	}
}

func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

func (p *Pool[T]) Put(x T) {
	p.pool.Put(x)
}
