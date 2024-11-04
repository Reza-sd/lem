package main

import (
	"fmt"
)

/*
The Object Pool design pattern is a creational design pattern that aims to reuse objects to improve performance and reduce object creation overhead. It's particularly useful for resource-intensive objects or those that are frequently created and destroyed.

Understanding the Pattern

- Object Pool: A collection of pre-initialized objects ready to be used.

- Acquisition: When an object is needed, it's acquired from the pool.

- Release: After use, the object is returned to the pool for future reuse.
*/
//============================================
//1. Define the Poolable Object Interface:

type Poolable interface {
	// Initialize the object for reuse
	Init()
}

//=====================================
//2. Create the Object Pool Structure:
type ObjectPool struct {
	pool    chan Poolable
	factory func() Poolable
}

//-------
func NewObjectPool(size uint8, theFactory func() Poolable) *ObjectPool {
	println("new objectpool")
	p := &ObjectPool{
		pool:    make(chan Poolable, size),
		factory: theFactory,
	}

	// Pre-populate the pool
	for i := 0; i < int(size); i++ {
		p.pool <- p.factory()
	}

	return p
}

//-------
func (p *ObjectPool) Acquire() (Poolable, error) {
	select {
	case obj := <-p.pool: // Try to get an object from the pool
		return obj, nil
	default:
		return nil, fmt.Errorf("error: no more object to Acquire") // Create a new object if pool is empty
	}
	// println("get the object")
	// obj := <-p.pool
	// obj.Init()
	// return obj
}

//-------
func (p *ObjectPool) Release(obj Poolable) {
	// Try to put the object back in the pool
	select {
	case p.pool <- obj:
		// Successfully returned to pool
		println("Release the object")
	default:
		// Pool is full, discard the object
		println("error: nothing to release")
	}
	// println("put the object")
	// p.pool <- obj
}

//============================================
//3. Using the Object Pool:
type MyObject struct {
	// ... object fields
}

func (m *MyObject) Init() {
	// Initialize the object here
	println("initialize the object")
}

func (m *MyObject) Do(n uint8) {
	// Initialize the object here
	println(n, " Do some jobs....")
}

//============================================
func main() {
	//---------------
	myPool := NewObjectPool(1, func() Poolable {
		return &MyObject{}
	})
	//---------------
	// Acquire objects from the pool
	obj1, err := myPool.Acquire() //.(*MyObject)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("obj1=%+v\n", obj1.(*MyObject))
		obj1.(*MyObject).Do(1)

	}
	//---------------
	// Use the object
	// ...
	//obj1.Do()
	obj2, err := myPool.Acquire() //.(*MyObject)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("obj2=%+v\n", obj2.(*MyObject))
		obj2.(*MyObject).Do(2)
	}
	// fmt.Printf("obj2=%+v\n",obj2)
	//---------------
	// // Return the object to the pool
	myPool.Release(obj1)
	//---------------
	obj3, err := myPool.Acquire() //.(*MyObject)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("obj3=%+v\n", obj3.(*MyObject))
		obj3.(*MyObject).Do(3)
	}
	//---------------
}

//============================================
