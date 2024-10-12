package main

import "fmt"

// ====================================
// Implementor interface
type Implementor interface {
	Operation() string
}

// ====================================
// Abstraction interface
type Abstraction interface {
	Operation() string
}

// ====================================
// RefinedAbstraction1 implements Abstraction
type RefinedAbstraction1 struct {
	implementor Implementor
}

func (ra1 *RefinedAbstraction1) Operation() string {
	return "RefinedAbstraction1: " + ra1.implementor.Operation()
}

// ====================================
// RefinedAbstraction2 implements Abstraction
type RefinedAbstraction2 struct {
	implementor Implementor
}

func (ra2 *RefinedAbstraction2) Operation() string {
	return "RefinedAbstraction2: " + ra2.implementor.Operation()
}

//====================================

// ConcreteImplementorA implements Implementor
type ConcreteImplementorA struct{}

func (cia *ConcreteImplementorA) Operation() string {
	return "ConcreteImplementorA"
}

// ConcreteImplementorB implements Implementor
type ConcreteImplementorB struct{}

func (cib *ConcreteImplementorB) Operation() string {
	return "ConcreteImplementorB"
}

// ====================================
func main() {
	// Create refined abstractions with different implementors
	ra1 := &RefinedAbstraction1{implementor: &ConcreteImplementorA{}}
	ra2 := &RefinedAbstraction2{implementor: &ConcreteImplementorB{}}

	// Use the refined abstractions
	fmt.Println(ra1.Operation()) // Output: RefinedAbstraction1: ConcreteImplementorA
	fmt.Println(ra2.Operation()) // Output: RefinedAbstraction2: ConcreteImplementorB
}

//====================================
