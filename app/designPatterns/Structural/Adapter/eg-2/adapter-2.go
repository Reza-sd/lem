package main

import "fmt"

//==============================================
/*
Adapter Pattern in Golang

The Adapter pattern is a structural design pattern that allows incompatible interfaces to work together. It provides a way to wrap an existing interface in a new interface that clients expect. This is useful when you need to integrate existing code with new systems or when you want to reuse existing classes without modifying their source code.

Key Components:

Target: Defines the interface that clients expect.
Adaptee: Defines the existing interface that needs to be adapted.
Adapter: Adapts the Adaptee interface to the Target interface.
*/
//==============================================
// Target interface defines the expected interface.
type Target interface {
	Request() string
}

//==============================================
// Adaptee defines the existing interface.
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter adapts the Adaptee interface to the Target interface.
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

//==============================================
func main() {
	// Create an Adaptee object.
	adaptee := &Adaptee{}

	// Create an Adapter object and wrap the Adaptee.
	adapter := &Adapter{adaptee: adaptee}

	// Use the Adapter to call the SpecificRequest method.
	result := adapter.Request()
	fmt.Println(result)
}

//==============================================
