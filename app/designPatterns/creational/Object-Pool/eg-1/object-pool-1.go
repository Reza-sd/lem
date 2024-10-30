package main

/*
The Object Pool design pattern is a creational design pattern that aims to reuse objects to improve performance and reduce object creation overhead. It's particularly useful for resource-intensive objects or those that are frequently created and destroyed.

Understanding the Pattern
Object Pool: A collection of pre-initialized objects ready to be used.
Acquisition: When an object is needed, it's acquired from the pool.
Release: After use, the object is returned to the pool for future reuse.
*/
//----------------------------------------
//1. Define the Poolable Object Interface:

type Poolable interface {
	// Initialize the object for reuse
	Init()
}

//----------------------------------------
//2. Create the Object Pool Structure:
type ObjectPool struct {
	pool    chan Poolable
	factory func() Poolable
}

func NewObjectPool(size int, factory func() Poolable) *ObjectPool {
	p := &ObjectPool{
		pool:    make(chan Poolable, size),
		factory: factory,
	}

	// Pre-populate the pool
	for i := 0; i < size; i++ {
		p.pool <- p.factory()
	}

	return p
}

func (p *ObjectPool) Get() Poolable {
	obj := <-p.pool
	obj.Init()
	return obj
}

func (p *ObjectPool) Put(obj Poolable) {
	p.pool <- obj
}

//----------------------------------------
//3. Using the Object Pool:
type MyObject struct {
	// ... object fields
}

func (m *MyObject) Init() {
	// Initialize the object here
}

//----------------------------------------
func main() {
	pool := NewObjectPool(10, func() Poolable {
		return &MyObject{}
	})

	// Acquire objects from the pool
	obj1 := pool.Get().(*MyObject)
	// Use the object
	// ...

	// Return the object to the pool
	pool.Put(obj1)
}
