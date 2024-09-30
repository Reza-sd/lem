package main

import "fmt"

//================================
/*
>>>>below is a simple implement of Singleton. in the case which there are no goroutines.<<<
*/
//================================
/*
Singleton Pattern in Golang

The Singleton pattern is a creational design pattern that ensures a class has only one instance and provides a global point of access to it. This is useful for scenarios where you need to control the number of instances of a class and share it across different parts of your application.
*/

//===================================
// Singleton class
type Singleton struct {
	data int
}

// Single instance (Global)
var instance *Singleton

// Private constructor
func newSingleton() *Singleton {
	return &Singleton{data: 0}
}

//===================================
// Public static method to access the single instance
func GetInstance() *Singleton {
	fmt.Println("instance=", instance)
	if instance == nil {
		instance = newSingleton()
	}
	return instance
}

//===================================
func main() {
	// Access the single instance
	singleton1 := GetInstance()
	singleton1.data = 10

	// Access the same instance from another part of the application
	singleton2 := GetInstance()
	fmt.Println(singleton2.data) // Output: 10
}

//===================================
