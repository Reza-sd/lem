package main

import "fmt"

// Define an interface
type Animal interface {
    Speak() string
}

// Define a struct implementing the interface
type Cat struct {
    Name string
}

// Implement the Speak method for Cat
func (c Cat) Speak() string {
    return "Meow!"
}

// Define another struct implementing the interface
type Dog struct {
    Name string
}

// Implement the Speak method for Dog
func (d Dog) Speak() string {
    return "Woof!"
}

// Define a Zoo struct that contains an Animal interface
type Zoo struct {
    animal Animal
}

func main() {
    // Create instances of Cat and Dog
    myCat := Cat{Name: "Whiskers"}
    myDog := Dog{Name: "Buddy"}

    // Create a Zoo instance with Cat
    zoo1 := Zoo{animal: myCat}
	
    fmt.Println(zoo1.animal.Speak()) // Output: Meow!

    // Create a Zoo instance with Dog
    zoo2 := Zoo{animal: myDog}
    fmt.Println(zoo2.animal.Speak()) // Output: Woof!
}
