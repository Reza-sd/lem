package main

import "fmt"

// ==============================================
// Implementor interface
type DrawAPI interface {
	DrawCircle(radius, x, y int)
}

// ==============================================
// Concrete Implementor 1
type RedCircle struct{}

func (r *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: red, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}

// ==============================================
// Concrete Implementor 2
type GreenCircle struct{}

func (g *GreenCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: green, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}

// ==============================================
// Abstraction
type Shape interface {
	Draw()
}

// Refined Abstraction
type Circle struct {
	x, y, radius int
	drawAPI      DrawAPI
}

func (c *Circle) Draw() {
	c.drawAPI.DrawCircle(c.radius, c.x, c.y)
}

// ==============================================
func main() {
	redCircle := &Circle{100, 100, 10, &RedCircle{}}
	greenCircle := &Circle{100, 100, 10, &GreenCircle{}}

	redCircle.Draw()
	greenCircle.Draw()
}

//==============================================
