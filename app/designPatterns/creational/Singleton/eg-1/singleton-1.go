package main

/*
The Singleton pattern in Golang ensures that a struct has only one instance and provides a global point of access to it. Here's an overview of implementing the Singleton pattern in Go:

*/

import (
	"fmt"
	"sync"
)

type Singleton struct {
	// fields
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

/*
Explanation
The Singleton struct is defined with its fields.
A private instance variable holds the single instance.
sync.Once is used to ensure thread-safe initialization.
GetInstance() is the public method to access the singleton instance.
*/
func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 == s2 {
		fmt.Println("Singleton works, both variables contain the same instance.")
	}
}
