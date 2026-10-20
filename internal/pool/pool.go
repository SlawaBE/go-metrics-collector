// Package pool provides a generic, type-safe object pool. Objects placed back
// into the pool are reset to their zero state before being reused, which makes
// the pool safe for the types implementing a Reset method.
package pool

import "sync"

// Resetter is the constraint imposed on the pooled type. A pooled value must be
// able to return itself to its zero state.
type Resetter interface {
	Reset()
}

// Pool is a generic container that reuses objects of a single concrete type.
// It mirrors sync.Pool but is type-safe and resets an object before it is
// returned to the pool, so callers always obtain a clean instance.
//
// The generic parameter T is constrained to types that implement Resetter. In
// this project those types declare Reset on a pointer receiver (for example
// *model.Metric), so Pool is typically instantiated with a pointer type and
// Get returns a fresh, reset *Metric.
type Pool[T Resetter] struct {
	pool *sync.Pool
}

// New creates and returns a Pool for the type T. The factory function
// allocates a fresh object whenever the pool is empty.
func New[T Resetter](factory func() T) *Pool[T] {
	return &Pool[T]{
		pool: &sync.Pool{
			New: func() any {
				return factory()
			},
		},
	}
}

// Get returns an object from the pool. If the pool is empty, a new object is
// created with the factory passed to New.
func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

// Put resets the object and returns it to the pool, making it safe for
// subsequent reuse by another caller.
func (p *Pool[T]) Put(obj T) {
	obj.Reset()
	p.pool.Put(obj)
}
