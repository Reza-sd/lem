package main

/*
The Open-Closed Principle (OCP) is one of the five SOLID principles of object-oriented design. It states that software entities (like classes, modules, functions, etc.) should be open for extension but closed for modification. In practice, this means that the behavior of a module should be extendable without modifying its source code.

Although Go (Golang) is not an object-oriented language in the traditional sense, you can still apply the OCP principle in Go using interfaces and composition. Here's how you can apply OCP in Golang:
*/
import "fmt"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	width, height float64
}

// Implementing Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Circle struct
type Circle struct {
	radius float64
}

// Implementing Shape interface for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// New function to calculate total area for any shape
func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	r := Rectangle{width: 5, height: 4}
	c := Circle{radius: 3}

	// Add new shapes to the slice
	shapes := []Shape{r, c}

	fmt.Println("Total Area:", TotalArea(shapes))
}
