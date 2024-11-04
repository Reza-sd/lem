package main

import (
	"fmt"
)

//=======================================
// 1- Prototype defines the interface that all prototypes must implement.
type Prototype interface {
	Clone() Prototype
	GetType() string
}

//=======================================
// 2- ConcretePrototype1 is a specific implementation of the Prototype interface.
type ConcretePrototype1 struct {
	Name string
}

//----------
// Clone creates a copy of ConcretePrototype1.
func (p *ConcretePrototype1) Clone() Prototype {
	return &ConcretePrototype1{Name: p.Name}
}

//----------
// GetType returns the type of the prototype (for demo purposes).
func (p *ConcretePrototype1) GetType() string {
	return "ConcretePrototype1"
}

//=======================================
// ConcretePrototype2 is another specific implementation of the Prototype interface.
type ConcretePrototype2 struct {
	Age int
}

//----------
// Clone creates a copy of ConcretePrototype2.
func (p *ConcretePrototype2) Clone() Prototype {
	return &ConcretePrototype2{Age: p.Age}
}

//----------
// GetType returns the type of the prototype.
func (p *ConcretePrototype2) GetType() string {
	return "ConcretePrototype2"
}

//=======================================
// PrototypeRegistry holds a collection of prototypes and allows cloning them.
type PrototypeRegistry struct {
	prototypes map[string]Prototype
}

//----------
// NewPrototypeRegistry initializes a new registry with an empty prototype map.
func NewPrototypeRegistry() *PrototypeRegistry {
	return &PrototypeRegistry{prototypes: make(map[string]Prototype)}
}

//----------
// RegisterPrototype adds a new prototype to the registry.
func (r *PrototypeRegistry) RegisterPrototype(key string, prototype Prototype) {
	r.prototypes[key] = prototype
}

//----------
// GetPrototype retrieves a clone of the prototype associated with the given key.
func (r *PrototypeRegistry) GetPrototype(key string) (Prototype, error) {
	if prototype, exists := r.prototypes[key]; exists {
		return prototype.Clone(), nil
	}
	return nil, fmt.Errorf("prototype not found for key: %s", key)
}

//=======================================
func main() {
	// Initialize the prototype registry.
	registry := NewPrototypeRegistry()

	// Register prototypes with unique keys.
	registry.RegisterPrototype("proto1", &ConcretePrototype1{Name: "Prototype 1"})
	registry.RegisterPrototype("proto2", &ConcretePrototype2{Age: 25})

	// Retrieve and clone prototypes from the registry.
	clone1, err := registry.GetPrototype("proto1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cloned %s: %+v\n", clone1.GetType(), clone1)
	}

	clone2, err := registry.GetPrototype("proto2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cloned %s: %+v\n", clone2.GetType(), clone2)
	}
}
