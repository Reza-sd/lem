package main

import (
	"fmt"
	"sync"
)

//---------------------------
/*
Explanation:

1-Object Structure: Represents the objects managed by the pool.

2- ObjectPool Structure: Manages a channel of objects. It uses a mutex to ensure thread safety when returning objects to the pool.

3- NewObjectPool: Initializes the pool with a specified number of objects.

4- Get: Retrieves an object from the pool. If the pool is empty, it returns nil.

5- Put: Returns an object to the pool. If the pool is full, the object is discarded.
*/
//------------------------------------
// Object represents a resource that needs to be managed by the pool
type Object struct {
	ID int
}

// ObjectPool manages a pool of reusable objects
type ObjectPool struct {
	pool  chan *Object
	mutex sync.Mutex
}

// NewObjectPool creates a new ObjectPool with a specified size
func NewObjectPool(size int) *ObjectPool {
	pool := make(chan *Object, size)
	for i := 0; i < size; i++ {
		pool <- &Object{ID: i}
	}
	return &ObjectPool{pool: pool}
}

// Get retrieves an object from the pool
func (p *ObjectPool) Get() *Object {
	select {
	case obj := <-p.pool:
		return obj
	default:
		return nil // or create a new object if preferred
	}
}

// Put returns an object to the pool
func (p *ObjectPool) Put(obj *Object) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	select {
	case p.pool <- obj:
	default:
		// Pool is full, discard the object
	}
}

// Example usage
func main() {
	pool := NewObjectPool(5)

	// Get an object from the pool
	obj := pool.Get()
	if obj != nil {
		fmt.Printf("Got object with ID: %d\n", obj.ID)
		// Use the object...

		// Return the object to the pool
		pool.Put(obj)
	} else {
		fmt.Println("No available objects in the pool")
	}
}
