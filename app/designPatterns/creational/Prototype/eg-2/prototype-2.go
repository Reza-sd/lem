package main

import "fmt"

//=========================================
/*
The Prototype pattern is a creational design pattern that provides a way to create new objects by copying existing objects. This is useful when creating complex objects with many attributes, as it avoids the need to re-create the entire object from scratch.
*/
//========================================
// Prototype interface defines the method for cloning objects.
type Prototype interface {
	Clone() Prototype
}

// ==============================================
// ConcreteProductA implements the Prototype interface.
type ConcreteProductA struct {
	Name string
}

func (a *ConcreteProductA) Clone() Prototype {
	return &ConcreteProductA{Name: a.Name}
}

// ---------------------
// ConcreteProductB implements the Prototype interface.
type ConcreteProductB struct {
	Name  string
	Value int
}

func (b *ConcreteProductB) Clone() Prototype {
	return &ConcreteProductB{Name: b.Name, Value: b.Value}
}

// ==============================================
type PrototypeRegistry struct {
	items map[string]Prototype
}

// NewPrototypeRegistry initializes a new registry with an empty prototype map.
func NewPrototypeRegistry() *PrototypeRegistry {
	return &PrototypeRegistry{items: make(map[string]Prototype)}
}

// RegisterPrototype adds a new prototype to the registry.
func (r *PrototypeRegistry) RegisterPrototype(key string, prototype Prototype) {
	r.items[key] = prototype
}

// GetPrototype retrieves a clone of the prototype associated with the given key.
func (r *PrototypeRegistry) GetPrototype(key string) (Prototype, error) {
	if prototype, exists := r.items[key]; exists {
		return prototype.Clone(), nil
	}
	return nil, fmt.Errorf("prototype not found for key: %s", key)
}

// ==============================================
func main() {
	// Create a prototype object.
	prototypeA := &ConcreteProductA{Name: "Product A"}

	// Clone the prototype to create a new object.

	cloneA := prototypeA.Clone()
	fmt.Println("Clone A:", cloneA)

	// Create another prototype object.
	prototypeB := &ConcreteProductB{Name: "Product B", Value: 10}

	// Clone the prototype to create a new object.
	cloneB := prototypeB.Clone()
	fmt.Println("Clone B:", cloneB)
	//--------------------------------
	// Initialize the prototype registry.
	myRegistry := NewPrototypeRegistry()
	// Register prototypes with unique keys.
	myRegistry.RegisterPrototype("A", &ConcreteProductA{Name: "Product A"})
	myRegistry.RegisterPrototype("B", &ConcreteProductB{Name: "Product B", Value: 10})

	// Retrieve and clone prototypes from the registry.
	/*
			%s

		This is a verb that formats the first argument as a string.
		It's commonly used to print plain text strings.
		%+v

		This is a verb with a flag +.
		The %v part formats the second argument in its default format.
		The + flag adds additional information, typically:
		For structs: It includes field names along with their values.
		For other types: It might provide more detailed output, such as pointers or addresses.
	*/

	clone1_A, err := myRegistry.GetPrototype("A")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cloned A %+v\n", clone1_A)
	}

	clone2_B, err := myRegistry.GetPrototype("B")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cloned B %+v\n", clone2_B)
	}
	//

	//------------------------------------
}

//==============================================
