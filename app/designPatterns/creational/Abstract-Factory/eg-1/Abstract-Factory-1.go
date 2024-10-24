package main

import "fmt"

/*
https://refactoring.guru/design-patterns/abstract-factory
- A factory is a class that returns products of a particular kind.

- The Abstract Factory pattern is a creational design pattern that provides an interface for creating families of related or dependent objects without specifying their concrete classes. It is useful when a system must work with multiple types of related objects, but you don't want the code to depend on their specific implementations.


*/

//==========================================
// Step 1: Define Product Interfaces
// We first define interfaces for the products. These are the Sedan and SUV products.

// Product A: Sedan interface
type Sedan interface {
	Drive()
}

// Product B: SUV interface
type SUV interface {
	Offroad()
}

//==========================================
//Step 2: Concrete Implementations of Products

// Concrete Product A1: Sporty Sedan
type SportySedan struct{}

func (s *SportySedan) Drive() {
	fmt.Println("Driving a sporty sedan!")
}

// Concrete Product A2: Luxury Sedan
type LuxurySedan struct{}

func (l *LuxurySedan) Drive() {
	fmt.Println("Driving a luxury sedan!")
}

//--------

// Concrete Product B1: Sporty SUV
type SportySUV struct{}

func (s *SportySUV) Offroad() {
	fmt.Println("Off-roading in a sporty SUV!")
}

// Concrete Product B2: Luxury SUV
type LuxurySUV struct{}

func (l *LuxurySUV) Offroad() {
	fmt.Println("Off-roading in a luxury SUV!")
}

//==========================================
// Step 3: Define the Abstract Factory Interface
// We then define the abstract factory interface that will be responsible for creating the products.

// Abstract Factory Interface
type CarFactory interface {
	MakeSedan() Sedan
	MakeSUV() SUV
}

//==========================================
//Step 4: Concrete Factories
//Now, let's implement the concrete factories that produce the actual cars (Sporty or Luxury).

// Concrete Factory 1: Sporty Car Factory
type SportyCarFactory struct{}

func (f *SportyCarFactory) MakeSedan() Sedan {
	return &SportySedan{}
}

func (f *SportyCarFactory) MakeSUV() SUV {
	return &SportySUV{}
}

// Concrete Factory 2: Luxury Car Factory
type LuxuryCarFactory struct{}

func (f *LuxuryCarFactory) MakeSedan() Sedan {
	return &LuxurySedan{}
}

func (f *LuxuryCarFactory) MakeSUV() SUV {
	return &LuxurySUV{}
}

// ---------------------------
func creatCars(thisFactory CarFactory) {
	Sedan := thisFactory.MakeSedan()
	SUV := thisFactory.MakeSUV()

	Sedan.Drive()
	SUV.Offroad()
}

// ==========================================
func main() {
	var factory CarFactory
	//----------------------
	// Using the SportyCarFactory
	factory = &SportyCarFactory{}
	sportySedan := factory.MakeSedan()
	sportySUV := factory.MakeSUV()

	sportySedan.Drive()
	sportySUV.Offroad()
	//----------------------
	// Using the LuxuryCarFactory
	factory = &LuxuryCarFactory{}
	luxurySedan := factory.MakeSedan()
	luxurySUV := factory.MakeSUV()

	luxurySedan.Drive()
	luxurySUV.Offroad()

	//----------------------
	//using creation function
	thisFactory := &LuxuryCarFactory{}
	creatCars(thisFactory)

}

//==========================================
