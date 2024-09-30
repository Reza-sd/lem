package main

import "fmt"

//======================================================
/*
Understanding the Composite Pattern:

Treats objects in a tree structure: The Composite Pattern allows you to treat individual objects and compositions of objects uniformly. This means that clients can interact with both simple objects and complex compositions of objects in the same way.

Core components:

Component: Defines the common interface for both leaf and composite objects.

Leaf: Represents the end nodes of the tree structure.

Composite: Represents the internal nodes of the tree structure, containing a list of child components.

Relationship: The Composite objects can contain both other Composite objects and Leaf objects, forming a hierarchical tree structure. This enables clients to treat all objects in the tree uniformly, regardless of their complexity.

*/
//=======================================================
// Component interface
type Component interface {
	Operation() string
}

//=======================================================
// Leaf implements Component
type Leaf struct {
	name string
}

func (l *Leaf) Operation() string {
	return fmt.Sprintf("Leaf: %s", l.name)
}

// Composite implements Component
type Composite struct {
	name     string
	children []Component
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Remove(component Component) {
	for i, child := range c.children {
		if child == component {
			c.children = append(c.children[:i], c.children[i+1:]...)
			break
		}
	}
}

func (c *Composite) GetChild(index int) Component {
	return c.children[index]
}

func (c *Composite) Operation() string {
	result := fmt.Sprintf("Composite: %s\n", c.name)
	for _, child := range c.children {
		result += child.Operation() + "\n"
	}
	return result
}

//=======================================================
func main() {
	// Create leaf objects
	leaf1 := &Leaf{name: "Leaf 1"}
	leaf2 := &Leaf{name: "Leaf 2"}

	// Create composite objects
	composite1 := &Composite{name: "Composite 1"}
	composite2 := &Composite{name: "Composite 2"}

	// Build the tree structure
	composite1.Add(leaf1)
	composite1.Add(composite2)
	composite2.Add(leaf2)

	// Use the composite objects
	fmt.Println(composite1.Operation())
}

//=======================================================
