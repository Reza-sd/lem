package main

import "fmt"
//--------------------------------------
// Product represents the final object being constructed.
type Product struct {
	PartA string
	PartB string
	PartC string
}
//--------------------------------------
// Builder interface defines the common methods for constructing the Product.
type Builder interface {
	BuildPartA(partA string) Builder
	BuildPartB(partB string) Builder
	BuildPartC(partC string) Builder

	GetProduct() Product
}
//--------------------------------------
// ConcreteBuilderA implements the Builder interface and constructs a Product with specific parts.
type ConcreteBuilderA struct {
	product *Product
}

func (b *ConcreteBuilderA) BuildPartA(partA string) Builder {
	b.product.PartA = partA
	return b
}

func (b *ConcreteBuilderA) BuildPartB(partB string) Builder {
	b.product.PartB = partB
	return b
}

func (b *ConcreteBuilderA) BuildPartC(partC string) Builder {
	b.product.PartC = partC
	return b
}

func (b *ConcreteBuilderA) GetProduct() Product {
	return *b.product
}
//---------------------------------------
type ConcreteBuilderB struct {
	product *Product
}

func (b *ConcreteBuilderB) BuildPartA(partA string) Builder {
	b.product.PartA = partA+" B"
	return b
}

func (b *ConcreteBuilderB) BuildPartB(partB string) Builder {
	b.product.PartB = partB+" B"
	return b
}

func (b *ConcreteBuilderB) BuildPartC(partC string) Builder {
	b.product.PartC = partC+" B"
	return b
}

func (b *ConcreteBuilderB) GetProduct() Product {
	return *b.product
}
//------------------------------------------
// Director orchestrates the construction process, calling methods on the builder.
type Director struct{}

func (d *Director) Construct(builder Builder) {
	builder.BuildPartA("Part A")
	builder.BuildPartB("Part B")
	builder.BuildPartC("Part C")
}
//---------------------------------------
func main() {
	// Create a concrete builder and use it to construct a Product.
	builder := &ConcreteBuilderA{product: &Product{}}
	director := &Director{}
	director.Construct(builder)

	product := builder.GetProduct()
	fmt.Println("Product:", product)

	//---------------------
	//OR
	builder2 :=&ConcreteBuilderB{product: &Product{}}
	director.Construct(builder2)
	product = builder.GetProduct()
	fmt.Println("Product:", product)
}
