package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

//--------------------------------
/*
https://golangbyexample.com/golang-object-pool/

The Object Pool Design Pattern is a creational design pattern where a pool of initialized objects is kept ready to use, rather than creating and destroying them on demand1
. This pattern is particularly useful when the cost of creating an object is high, and you need to manage a limited number of instances
*/
//--------------------------------
//1- Define the Pool Object Interface: This interface will be used to manage the pool objects.
type iPoolObject interface {
	//
	getID() string //This is any id which can be used to compare two different pool objects
}

//--------------------------------
//2- Create the Pool Structure: This structure will manage the pool of objects
type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	mulock   *sync.Mutex
}

//--------------------------------
//3-Initialize the Pool: Create a function to initialize the pool with a given number of objects
func initPool(poolObjects []iPoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("cannot create a pool of 0 length")
	}
	active := make([]iPoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}
	return pool, nil
}

//--------------------------------
//4-Loan an Object: Create a function to loan an object from the pool

func (p *pool) loan() (iPoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("no pool object free. Please request after sometime")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())

	return obj, nil
}

//--------------------------------
// 5- remove an Object:
func (p *pool) remove(target iPoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("targe pool object doesn't belong to the pool")
}

//--------------------------------
//6- Return an Object: Create a function to return an object to the pool
func (p *pool) receive(target iPoolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
	return nil
}

//==================================
//use
type connection struct {
	id string
}

func (c *connection) getID() string {
	return c.id
}

//--------------------------------
func main() {
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := initPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
	conn1, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	conn2, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	pool.receive(conn1)
	pool.receive(conn2)
}

//==================================
