package main

import "fmt"

//=========================================
/*
The Adapter Pattern is a structural design pattern that allows objects with incompatible interfaces to work together. In Golang, this pattern is particularly useful for integrating different interfaces without altering their existing code.

Here’s a simple example to illustrate the Adapter Pattern in Golang:

Define the Target Interface: This is the interface your client expects to work with.
Create the Adaptee: This is the existing interface that needs to be adapted.
Implement the Adapter: This wraps the Adaptee and makes it compatible with the Target Interface.
*/
//=========================================
// Target interface
type Target interface {
	Request() string
}

//=========================================
// Adaptee with an incompatible interface
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee's specific request"
}

// Adapter makes Adaptee compatible with Target
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

//=========================================
func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee}
	fmt.Println(adapter.Request())
}

//=========================================
