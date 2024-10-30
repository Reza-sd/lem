package main

import (
	"fmt"
	"sync"
)

// ObjectPool holds a pool of objects
type ObjectPool struct {
	m       sync.Mutex
	objects chan interface{}
	factory func() interface{}
}

// NewObjectPool creates a new object pool
func NewObjectPool(capacity int, factory func() interface{}) *ObjectPool {
	return &ObjectPool{
		objects: make(chan interface{}, capacity),
		factory: factory,
	}
}

// Get retrieves an object from the pool
func (p *ObjectPool) Get() interface{} {
	p.m.Lock()
	defer p.m.Unlock()

	select {
	case obj := <-p.objects:
		return obj
	default:
		return p.factory()
	}
}

// Put returns an object to the pool
func (p *ObjectPool) Put(obj interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	select {
	case p.objects <- obj:
	default:
	}
}

func main() {
	pool := NewObjectPool(5, func() interface{} {
		return &MyObject{ID: 0}
	})

	// Get an object from the pool
	obj1 := pool.Get().(*MyObject)
	fmt.Println("Object 1:", obj1)

	// Return the object to the pool
	pool.Put(obj1)

	// Get another object from the pool
	obj2 := pool.Get().(*MyObject)
	fmt.Println("Object 2:", obj2)
}

// MyObject is a sample object
type MyObject struct {
	ID int
}
