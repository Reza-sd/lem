package main

import "fmt"

/*
 Use the Abstract Factory when your code needs to work with various families of related products, but you don’t want it to depend on the concrete classes of those products—they might be unknown beforehand or you simply want to allow for future extensibility.


Let's implement another example of the Abstract Factory pattern in Golang. This time, let's work with a different domain: a Furniture Factory that produces two families of products — Victorian and Modern furniture. The furniture types will be Chair and Sofa.
*/
//==============================================
//Step 1: Define Product Interfaces

// Chair interface (Product A)
type Chair interface {
	SitOn()
}

// Sofa interface (Product B)
type Sofa interface {
	LieOn()
}

//==============================================
//Step 2: Concrete Implementations of Products

//Next, we create the concrete implementations for both Victorian and Modern families of chairs and sofas.

// Victorian Chair (Concrete Product A1)
type VictorianChair struct{}

func (vc *VictorianChair) SitOn() {
	fmt.Println("Sitting on a Victorian chair.")
}

// Victorian Sofa (Concrete Product B1)
type VictorianSofa struct{}

func (vs *VictorianSofa) LieOn() {
	fmt.Println("Lying on a Victorian sofa.")
}

// Modern Chair (Concrete Product A2)
type ModernChair struct{}

func (mc *ModernChair) SitOn() {
	fmt.Println("Sitting on a modern chair.")
}

// Modern Sofa (Concrete Product B2)
type ModernSofa struct{}

func (ms *ModernSofa) LieOn() {
	fmt.Println("Lying on a modern sofa.")
}

// ==============================================
// Step 3: Define the Abstract Factory Interface
// FurnitureFactory interface (Abstract Factory)
// define behaviour instead of concrete implementation
type FurnitureFactory interface {
	MakeChair() Chair
	MakeSofa() Sofa
}

// ==============================================
// Step 4: Concrete Factories
// Victorian Furniture Factory (Concrete Factory 1)
type VictorianFurnitureFactory struct{}

func (vf *VictorianFurnitureFactory) MakeChair() Chair {
	return &VictorianChair{}
}

func (vf *VictorianFurnitureFactory) MakeSofa() Sofa {
	return &VictorianSofa{}
}

// Modern Furniture Factory (Concrete Factory 2)
type ModernFurnitureFactory struct{}

func (mf *ModernFurnitureFactory) MakeChair() Chair {
	return &ModernChair{}
}

func (mf *ModernFurnitureFactory) MakeSofa() Sofa {
	return &ModernSofa{}
}

// ==============================================
// /Step 5: Client Code
// Finally, the client can work with the abstract factory to create different types of furniture without knowing the concrete implementations.
// Client code
func main() {
	var factory FurnitureFactory

	// Using the Victorian Furniture Factory
	factory = &VictorianFurnitureFactory{}
	victorianChair := factory.MakeChair()
	victorianSofa := factory.MakeSofa()

	victorianChair.SitOn()
	victorianSofa.LieOn()

	// Using the Modern Furniture Factory
	factory = &ModernFurnitureFactory{}
	modernChair := factory.MakeChair()
	modernSofa := factory.MakeSofa()

	modernChair.SitOn()
	modernSofa.LieOn()
}

//==============================================
