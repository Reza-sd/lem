package main

import "fmt"

//==============================================
/*
 The Composite pattern is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects.
*/
//==============================================
// Component interface
type Component interface {
	Operation() string
}

// ==============================================
// Leaf
type Leaf struct {
	name string
}

func (l *Leaf) Operation() string {
	return l.name
}

// ==============================================
// Composite
type Composite struct {
	name       string
	components []Component
}

func (c *Composite) Operation() string {
	result := c.name + " ["
	for i, component := range c.components {
		result += component.Operation()
		if i < len(c.components)-1 {
			result += " + "
		}
	}
	result += "]"
	return result
}

// ---------------------------------
func (c *Composite) Add(component Component) {
	c.components = append(c.components, component)
}

func (c *Composite) Remove(component Component) {
	for i, comp := range c.components {
		if comp == component {
			c.components = append(c.components[:i], c.components[i+1:]...)
			break
		}
	}
}

// ==============================================
func main() {
	leaf1 := &Leaf{name: "Leaf 1"}
	leaf2 := &Leaf{name: "Leaf 2"}
	leaf3 := &Leaf{name: "Leaf 3"}

	composite1 := &Composite{name: "Composite 1"}
	composite1.Add(leaf1)
	composite1.Add(leaf2)

	composite2 := &Composite{name: "Composite 2"}
	composite2.Add(leaf3)
	composite2.Add(composite1)

	fmt.Println(composite2.Operation())
}

//==============================================
