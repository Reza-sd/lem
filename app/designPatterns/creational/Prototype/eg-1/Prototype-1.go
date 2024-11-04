package main

import "fmt"

/*
The Prototype Pattern is a creational design pattern that allows you to create new objects by copying an existing object, known as a "prototype." This approach is useful when the cost of creating a new object is high (e.g., complex initialization), or when you want to avoid subclassing and prefer to dynamically decide what kind of object to create.



Example: Shapes (Circle and Rectangle)
We will create an example where we define a ShapePrototype interface, and two types of shapes: Circle and Rectangle. The Clone method will allow us to create a copy of these shapes.

*/
//==============================================
//1. Define the Prototype Interface (ShapePrototype):
// ShapePrototype is the prototype interface
type ShapePrototype interface {
	Clone() ShapePrototype
	GetDetails() string
}

// ==============================================
// 2. Concrete Prototypes (Circle and Rectangle):
// Circle is a concrete prototype
type Circle struct {
	Radius int
}

func (c *Circle) Clone() ShapePrototype {
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

func (r *Rectangle) Clone() ShapePrototype {
	// Return a copy of the rectangle
	return &Rectangle{Width: r.Width, Height: r.Height}
}

func (r *Rectangle) GetDetails() string {
	return fmt.Sprintf("Rectangle with width %d and height %d", r.Width, r.Height)
}

// ==============================================
// PrototypeRegistry holds a collection of prototypes and allows cloning them.

type PrototypeRegistry struct {
	prototypes map[string]ShapePrototype
}

// NewPrototypeRegistry initializes a new registry with an empty prototype map.
func NewPrototypeRegistry() *PrototypeRegistry {
	return &PrototypeRegistry{prototypes: make(map[string]ShapePrototype)}
}

// RegisterPrototype adds a new prototype to the registry.
func (r *PrototypeRegistry) RegisterPrototype(key string, prototype ShapePrototype) {
	r.prototypes[key] = prototype
}

// GetPrototype retrieves a clone of the prototype associated with the given key.
func (r *PrototypeRegistry) GetPrototype(key string) (ShapePrototype, error) {
	if prototype, exists := r.prototypes[key]; exists {
		return prototype.Clone(), nil
	}
	return nil, fmt.Errorf("prototype not found for key: %s", key)
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
	//-------------------------------------------------
	// Initialize the prototype registry.
	myRegistry := NewPrototypeRegistry()
	// Register prototypes with unique keys.
	myRegistry.RegisterPrototype("circleBig", &Circle{Radius: 999})
	myRegistry.RegisterPrototype("rectangleBig", &Rectangle{Width: 888, Height: 777})

	// Retrieve and clone prototypes from the registry.
	clone1_circle, err := myRegistry.GetPrototype("circleBig")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cloned %s: %+v\n", clone1_circle.GetDetails(), clone1_circle)
	}

	clone2_rectangle, err := myRegistry.GetPrototype("rectangleBig")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cloned %s: %+v\n", clone2_rectangle.GetDetails(), clone2_rectangle)
	}
	//--------------------------------------------------
}

//==============================================
