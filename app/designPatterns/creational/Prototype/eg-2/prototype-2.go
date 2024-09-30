package main

import "fmt"

//=========================================
/*
The Prototype pattern is a creational design pattern that provides a way to create new objects by copying existing objects. This is useful when creating complex objects with many attributes, as it avoids the need to re-create the entire object from scratch.
*/
//========================================
// Prototype interface defines the method for cloning objects.
type Prototype interface {
    Clone() Prototype
}
//==============================================
// ConcreteProductA implements the Prototype interface.
type ConcreteProductA struct {
    Name string
}

func (a *ConcreteProductA) Clone() Prototype {
    return &ConcreteProductA{Name: a.Name}
}
//---------------------
// ConcreteProductB implements the Prototype interface.
type ConcreteProductB struct {
    Name string
    Value int
}

func (b *ConcreteProductB) Clone() Prototype {
    return &ConcreteProductB{Name: b.Name, Value: b.Value}
}
//==============================================
func main() {
    // Create a prototype object.
    prototypeA := &ConcreteProductA{Name: "Product A"}

    // Clone the prototype to create a new object.
    cloneA := prototypeA.Clone().(*ConcreteProductA)
    fmt.Println("Clone A:", cloneA)

    // Create another prototype object.
    prototypeB := &ConcreteProductB{Name: "Product B", Value: 10}

    // Clone the prototype to create a new object.
    cloneB := prototypeB.Clone().(*ConcreteProductB)
    fmt.Println("Clone B:", cloneB)
}
//==============================================