package main

import (
	"fmt"
)

//==================================
/*
The Object Pool design pattern is a creational design pattern that manages a pool of objects, reusing them to avoid the overhead of creating and destroying objects repeatedly.

This is particularly useful for resource-intensive objects like database connections, network sockets, or thread pools.
*/
//====================================
type Object interface {
	// Initialize the object for reuse
	Init()
}

//====================================
//Define the Object Pool Structure:
//pool: A channel to hold available objects.
//factory: A function to create new objects when the pool is empty.

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
//Get an Object from the Pool:
//Tries to get an object from the pool.
//If the pool is empty, return an error.

func (p *ObjectPool) Acquire() (Object, error) {
	select {
	case obj := <-p.pool:
		//fmt.Println("Acquire:",obj)
		return obj, nil
	default:
		return nil, fmt.Errorf("error: no more connection to Acquire")
	}
}

//--------------
//Concurrency Safety: The select statement in Get() ensures thread-safe access to the pool.
//Select statements in Go are inherently thread-safe.

/*
However, it's important to note that while select statements are thread-safe, the code within each case block might not be. 

To ensure thread safety within a case block, you should use appropriate synchronization mechanisms like mutexes or channels to protect shared resources.

*/
//----------------
//Return an Object to the Pool:

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
