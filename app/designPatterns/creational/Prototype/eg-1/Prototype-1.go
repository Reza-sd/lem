package main

import "fmt"

/*
The Prototype Pattern is a creational design pattern that allows you to create new objects by copying an existing object, known as a "prototype." This approach is useful when the cost of creating a new object is high (e.g., complex initialization), or when you want to avoid subclassing and prefer to dynamically decide what kind of object to create.



Example: Shapes (Circle and Rectangle)
We will create an example where we define a Shape interface, and two types of shapes: Circle and Rectangle. The Clone method will allow us to create a copy of these shapes.

*/
//==============================================
//1. Define the Prototype Interface (Shape):
// Shape is the prototype interface
type Shape interface {
	Clone() Shape
	GetDetails() string
}

// ==============================================
// 2. Concrete Prototypes (Circle and Rectangle):
// Circle is a concrete prototype
type Circle struct {
	Radius int
}

func (c *Circle) Clone() Shape {
	// Return a copy of the circle
	return &Circle{Radius: c.Radius}
}

func (c *Circle) GetDetails() string {
	return fmt.Sprintf("Circle with radius %d", c.Radius)
}

// Rectangle is another concrete prototype
type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Clone() Shape {
	// Return a copy of the rectangle
	return &Rectangle{Width: r.Width, Height: r.Height}
}

func (r *Rectangle) GetDetails() string {
	return fmt.Sprintf("Rectangle with width %d and height %d", r.Width, r.Height)
}

// ==============================================
func main() {
	// Create a circle and clone it
	circle := &Circle{Radius: 10}
	clonedCircle := circle.Clone()

	fmt.Println(circle.GetDetails())       // Output: Circle with radius 10
	fmt.Println(clonedCircle.GetDetails()) // Output: Circle with radius 10

	// Create a rectangle and clone it
	rectangle := &Rectangle{Width: 20, Height: 10}
	clonedRectangle := rectangle.Clone()

	fmt.Println(rectangle.GetDetails())       // Output: Rectangle with width 20 and height 10
	fmt.Println(clonedRectangle.GetDetails()) // Output: Rectangle with width 20 and height 10
}

//==============================================
