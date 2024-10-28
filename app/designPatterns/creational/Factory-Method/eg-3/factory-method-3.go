package main

import "fmt"
//-------------------------------------------
// Product interface defines the common behavior for objects to be created.
type Product interface {
	DoSomething()
}
//-------------------------------------------
// ConcreteProductA implements the Product interface.
type ConcreteProductA struct{}

func (a *ConcreteProductA) DoSomething() {
	fmt.Println("ConcreteProductA: Doing something.")
}

// ConcreteProductB implements the Product interface.
type ConcreteProductB struct{}

func (b *ConcreteProductB) DoSomething() {
	fmt.Println("ConcreteProductB: Doing something.")
}
//-------------------------------------------
// Factory interface defines the method for creating Product instances.
type Factory interface {
	CreateProduct() Product
}
//-------------------------------------------
// ConcreteFactoryA implements the Factory interface and creates ConcreteProductA instances.
type ConcreteFactoryA struct{}

func (a *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteFactoryB implements the Factory interface and creates ConcreteProductB instances.
type ConcreteFactoryB struct{}

func (b *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}
//-------------------------------------------
func main() {
	// Use ConcreteFactoryA to create a ConcreteProductA instance.
	factoryA := &ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	productA.DoSomething()

	// Use ConcreteFactoryB to create a ConcreteProductB instance.
	factoryB := &ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	productB.DoSomething()
}
