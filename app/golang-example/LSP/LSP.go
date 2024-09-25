package main

/*
1. Liskov Substitution Principle (LSP)
The Liskov Substitution Principle is one of the five SOLID principles of object-oriented programming. It states:

"Objects of a superclass should be replaceable with objects of its subclasses without affecting the correctness of the program."

In simpler terms, if you have a function that accepts a base type (interface or struct in Go), you should be able to pass any derived type (implementations) without the function misbehaving. In Go, this typically applies when working with interfaces and structs.
*/
import "fmt"

// Animal is an interface with a Speak method.
type Animal interface {
	Speak() string
}

// Dog is a struct that implements the Animal interface.
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat is another struct that implements the Animal interface.
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Function that accepts any Animal
func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	// Both can be passed to the function because they implement the Animal interface
	MakeAnimalSpeak(dog)
	MakeAnimalSpeak(cat)
}
