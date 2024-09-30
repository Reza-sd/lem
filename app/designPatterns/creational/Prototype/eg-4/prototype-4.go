package main

import "fmt"

//==============================================
//Define a Cloneable interface:
type Cloneable interface {
	Clone() Cloneable
}

//==============================================
//Create concrete types that implement the Cloneable interface:
type ConcretePrototype struct {
	name string
}

func (p *ConcretePrototype) Clone() Cloneable {
	return &ConcretePrototype{name: p.name + "_clone"}
}

//==============================================
func main() {
	original := &ConcretePrototype{name: "Original"}
	clone := original.Clone().(*ConcretePrototype)
	clone2 := clone.Clone().(*ConcretePrototype)

	fmt.Println(original.name) // Output: Original
	fmt.Println(clone.name)    // Output: Original_clone
	fmt.Println(clone2.name)
}

//==============================================
