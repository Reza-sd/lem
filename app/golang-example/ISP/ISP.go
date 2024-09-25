package main

import "math"

/*
Improved code organization: Breaking down interfaces into smaller, more focused ones leads to a more organized and understandable code structure.
Enhanced reusability: Smaller, more specific interfaces are more likely to be reused in different parts of your application, reducing code duplication.
Increased maintainability: Changes to an interface are less likely to have far-reaching consequences when it's narrowly defined.
Better testability: Smaller interfaces are easier to test in isolation, leading to more robust unit tests.
Implementing ISP in Golang:

Identify common interface elements: Look for common methods or attributes that are used by multiple clients.
Create specialized interfaces: Create separate interfaces for each group of related methods or attributes.
Implement specialized interfaces: Implement the specialized interfaces by your structs or types.
Use composition: If a struct needs to implement multiple interfaces, use composition to avoid interface explosion.
*/
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Specialized interface for shapes with a perimeter
type Perimeterable interface {
	Perimeter() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func main() {

}
