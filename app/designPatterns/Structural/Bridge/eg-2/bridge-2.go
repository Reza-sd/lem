package main

import "fmt"

//=======================================
/*
The Bridge Pattern is a structural design pattern that separates an abstraction from its implementation, allowing them to vary independently. In Golang, it's particularly useful for decoupling interfaces from their implementations.

Key Components:

- Abstraction: Defines the high-level interface that client code uses.
- Refined Abstraction: Extends the Abstraction interface.
- Implementor: Defines the interface for implementation classes.
- Concrete Implementor: Implements the Implementor interface.
*/

//=======================================
// Implementor
type Renderer interface {
	RenderCircle(radius float64)
	RenderSquare(side float64)
}

//---------------------------
// Concrete Implementors
type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing a circle of radius %.2f as vectors\n", radius)
}

func (v *VectorRenderer) RenderSquare(side float64) {
	fmt.Printf("Drawing a square of side %.2f as vectors\n", side)
}

//---------------------------
type RasterRenderer struct{}

func (r *RasterRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing a circle of radius %.2f as pixels\n", radius)
}

func (r *RasterRenderer) RenderSquare(side float64) {
	fmt.Printf("Drawing a square of side %.2f as pixels\n", side)
}

//---------------------------
// Abstraction
type Shape interface {
	Draw()
}

//---------------------------
// Refined Abstractions
type Circle struct {
	renderer Renderer
	radius   float64
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

//---------------------------
type Square struct {
	renderer Renderer
	side     float64
}

func (s *Square) Draw() {
	s.renderer.RenderSquare(s.side)
}

//=======================================
// Client code
func main() {
	vectorRenderer := &VectorRenderer{}
	rasterRenderer := &RasterRenderer{}

	circle := &Circle{renderer: vectorRenderer, radius: 5}
	square := &Square{renderer: rasterRenderer, side: 4}

	circle.Draw()
	square.Draw()
}

//=======================================
