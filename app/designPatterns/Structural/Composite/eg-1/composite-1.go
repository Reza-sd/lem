package main

import "fmt"

//=======================================
/*
The Composite Pattern is a structural design pattern that allows you to compose objects into tree structures to represent part-whole hierarchies. This pattern lets clients treat individual objects and compositions of objects uniformly.
*/
//=========================================
// Component interface
type Graphic interface {
	Draw()
}

// =========================================
// Leaf
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing a Circle")
}

// Leaf
type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Drawing a Square")
}

// Composite
type CompositeGraphic struct {
	children []Graphic
}

func (cg *CompositeGraphic) Draw() {
	for _, child := range cg.children {
		child.Draw()
	}
}

func (cg *CompositeGraphic) Add(graphic Graphic) {
	cg.children = append(cg.children, graphic)
}

func (cg *CompositeGraphic) Remove(graphic Graphic) {
	for i, child := range cg.children {
		if child == graphic {
			cg.children = append(cg.children[:i], cg.children[i+1:]...)
			break
		}
	}
}

//========================================
/*
In this example:

Graphic is the component interface.
Circle and Square are leaf nodes that implement the Graphic interface.

CompositeGraphic is the composite node that can contain children of type Graphic.

This setup allows you to build complex structures by combining simple objects and other composite objects, and treat them uniformly.
*/
//=========================================
func main() {
	circle := &Circle{}
	square := &Square{}

	composite := &CompositeGraphic{}
	composite.Add(circle)
	composite.Add(square)

	composite.Draw()
}

//=========================================
