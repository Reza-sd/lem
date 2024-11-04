package main

import (
	"fmt"
)

//====================================
type Object interface {
	// Initialize the object for reuse
	Init()
}

//====================================
type ObjectPool struct {
	pool    chan Object
	factory func() Object
}

//----------------
func NewObjectPool(size int, factory func() Object) *ObjectPool {
	pool := make(chan Object, size)
	for i := 0; i < size; i++ {
		pool <- factory()
	}
	println("new connection pool, size:", size)
	return &ObjectPool{pool, factory}
}

//----------------
func (p *ObjectPool) Acquire() (Object, error) {
	select {
	case obj := <-p.pool:
		//fmt.Println("Acquire:",obj)
		return obj, nil
	default:
		return nil, fmt.Errorf("error: no more connection to Acquire")
	}
}

//----------------
func (p *ObjectPool) Release(obj Object) {
	// Try to put the object back in the pool
	select {
	case p.pool <- obj:
		// Successfully returned to pool
		println("Release: the connection")
	default:
		// Pool is full, discard the object
		println("error: nothing to release")
	}
}

//====================================
type DatabaseConnection struct {
	// ...
}

func (d *DatabaseConnection) Init() {
	println("initialize: the db connection")
}

//====================================
func main() {
	pool := NewObjectPool(1, func() Object {
		return new(DatabaseConnection)
	})

	// Get a connection from the pool
	conn1, err := pool.Acquire()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Acquire: conn1=%+v\n", conn1.(*DatabaseConnection))
	}

	conn2, err := pool.Acquire()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Acquire: conn1=%+v\n", conn2.(*DatabaseConnection))
	}

	// Return the connection to the pool
	pool.Release(conn1)

	conn3, err := pool.Acquire()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Acquire: conn3=%+v\n", conn3.(*DatabaseConnection))
	}
}

//====================================
