package main

import (
	"fmt"
	"sync"
)

//-------------------------------------
/*
Key Components

1- Object: Represents the resource being managed by the pool.

2- ObjectPool: Manages the pool of objects, handling their creation and reuse.

3- Get: Retrieves an object from the pool or creates a new one if necessary.

4- Put: Returns an object to the pool for future reuse.
*/

//----------------------------------------
// Object represents the expensive-to-create resource
type Object struct {
	ID int
}

// ObjectPool manages a pool of reusable objects
type ObjectPool struct {
	pool   chan *Object
	mu     sync.Mutex
	nextID int
}

// NewObjectPool initializes the ObjectPool with a given size
func NewObjectPool(size int) *ObjectPool {
	return &ObjectPool{
		pool: make(chan *Object, size),
	}
}

// Get provides a reusable object or creates a new one if the pool is empty
func (p *ObjectPool) Get() *Object {
	p.mu.Lock()
	defer p.mu.Unlock()

	select {
	case obj := <-p.pool:
		return obj
	default:
		p.nextID++
		return &Object{ID: p.nextID}
	}
}

// Put returns an object back to the pool
func (p *ObjectPool) Put(obj *Object) {
	select {
	case p.pool <- obj:
		// Object successfully returned to the pool
	default:
		// Pool is full, discard the object
	}
}

func main() {
	pool := NewObjectPool(2)

	obj1 := pool.Get()
	fmt.Printf("Got object with ID: %d\n", obj1.ID)

	obj2 := pool.Get()
	fmt.Printf("Got object with ID: %d\n", obj2.ID)

	// Returning objects to the pool
	pool.Put(obj1)
	pool.Put(obj2)

	obj3 := pool.Get()
	fmt.Printf("Got object with ID: %d\n", obj3.ID)
}
