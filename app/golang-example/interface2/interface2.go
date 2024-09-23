package main

import "fmt"

// Define an interface with two method signatures
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define Rectangle type
type Rectangle struct {
	Width, Height float64
}

// Implement the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Define Circle type
type Circle struct {
	Radius float64
}

// Implement the Shape interface for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}
func main() {
	var s Shape

	s = Rectangle{Width: 5, Height: 7}
	fmt.Printf("Rectangle Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())

	s = Circle{Radius: 3.5}
	fmt.Printf("Circle Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())

	r := Rectangle{Width: 5, Height: 7}
	fmt.Printf("Rectangle Area: %f, Perimeter: %f\n", r.Area(), r.Perimeter())

	c := Circle{Radius: 3.5}
	fmt.Printf("Circle Area: %f, Perimeter: %f\n", c.Area(), c.Perimeter())
}
