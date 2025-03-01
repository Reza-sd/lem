package main

import "fmt"

type BrandType uint

const (
	adidasType BrandType = iota
	nikeType
)

/*
Abstract Factory Design Pattern is a creational design pattern that lets you create a family of related objects. It is an abstraction over the factory pattern. It is best explained with an example. Let’s say we have two factories

nike
adidas
Imagine you need to buy a sports kit which has a shoe and short. Preferably most of the time you would want to buy a full sports kit of a similar factory i.e either nike or adidas. This is where the abstract factory comes into the picture as concrete products that you want is shoe and a short and these products will be created by the abstract factory of nike and adidas.
Both these two factories – nike and adidas implement iSportsFactory interface.
We have two product interfaces.

iShoe – this interface is implemented by nikeShoe and adidasShoe concrete product.
iShort – this interface is implemented by nikeShort and adidasShort concrete product.
*/
//=============================================
//Step 1: Define Product Interfaces
//product A
type iShort interface {
	setLogo(logo string)
	setSize(size uint8)
	getLogo() string
	getSize() uint8
}

//-------------------
//product B
type iShoe interface {
	setLogo(logo string)
	setSize(size uint8)
	getLogo() string
	getSize() uint8
}

//=============================================
//Step 2: Concrete Implementations of Products
//product A
type shoe struct {
	logo string
	size uint8
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size uint8) {
	s.size = size
}

func (s *shoe) getSize() uint8 {
	return s.size
}

//---------------------
//product B
type short struct {
	logo string
	size uint8
}

func (s *short) setLogo(logo string) {
	s.logo = logo
}

func (s *short) getLogo() string {
	return s.logo
}

func (s *short) setSize(size uint8) {
	s.size = size
}

func (s *short) getSize() uint8 {
	return s.size
}

//=============================================
// Step 3: Define the Abstract Factory Interface
type iSportsFactory interface {
	makeShoe(uint8) iShoe
	makeShort(uint8) iShort
}

//=============================================
// Step 4: Concrete Factories
//Factory 1(adidas):
type adidasFactory struct {
}

func (a *adidasFactory) makeShoe(theSize uint8) iShoe {
	return &shoe{
		logo: "adidas",
		size: theSize,
	}
}

func (a *adidasFactory) makeShort(theSize uint8) iShort {
	return &short{
		logo: "adidas",
		size: theSize,
	}
}

//------------------
//Factory 2(nike):

type nikeFactory struct {
}

func (n *nikeFactory) makeShoe(theSize uint8) iShoe {
	return &shoe{
		logo: "nike",
		size: theSize,
	}
}

func (n *nikeFactory) makeShort(theSize uint8) iShort {
	return &short{
		logo: "nike",
		size: theSize,
	}
}

//=============================================

func getSportsFactory(brand BrandType) (iSportsFactory, error) {
	if brand == adidasType {
		return &adidasFactory{}, nil
	}
	if brand == nikeType {
		return &nikeFactory{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}

//=============================================
func main() {
	adidasFactory, _ := getSportsFactory(adidasType)
	nikeFactory, _ := getSportsFactory(nikeType)
	nikeShoe := nikeFactory.makeShoe(30)
	nikeShort := nikeFactory.makeShort(45)
	adidasShoe := adidasFactory.makeShoe(32)
	adidasShort := adidasFactory.makeShort(41)

	fmt.Println(adidasShoe.getLogo(), adidasShoe.getSize())
	fmt.Println(adidasShort.getLogo(), adidasShort.getSize())

	fmt.Println(nikeShoe.getLogo(), nikeShoe.getSize())
	fmt.Println(nikeShort.getLogo(), nikeShort.getSize())

}

//=============================================
