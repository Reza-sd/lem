package main

import "fmt"

/*
The Factory Method pattern is a creational design pattern that provides an interface for creating objects in a superclass but allows subclasses to alter the type of objects that will be created. In Go (Golang), you can implement the Factory Method pattern to provide flexibility in object creation without exposing the instantiation logic to the client.

1. Basic Structure
- Product Interface: Defines the interface for objects the factory method will create.
- Concrete Products: Concrete implementations of the product interface.
- Creator Interface: Declares the factory method that returns an object of the product type.
- Concrete Creators: Override the factory method to produce specific types of products.


//-----------
Components:
1- Product Interface: Defines the interface that all concrete products must implement.
2- Concrete Products: Specific implementations of the product interface.
3- Creator Interface: Declares the factory method for creating products.
4- Concrete Creators: Implement the factory method to create specific products.

//-----------
*/
//==============================================
//1. Define the Product Interface (Transport):
// Transport is the product interface
type Transport interface {
	Deliver() string
}

// ==============================================
// 2. Concrete Products (Car and Bike):
// Car is a concrete product implementing the Transport interface
type Car struct{}

func (c *Car) Deliver() string {
	return "Delivering by car"
}

// Bike is another concrete product implementing the Transport interface
type Bike struct{}

func (b *Bike) Deliver() string {
	return "Delivering by bike"
}

// ==============================================
// 3. Define the Creator Interface:
// TransportFactory is the creator interface that declares the factory method.
type TransportFactory interface {
	CreateTransport() Transport
}

// ==============================================
// 4. Concrete Creators (CarFactory and BikeFactory)
// CarFactory is a concrete creator that returns a Car.
type CarFactory struct{}

func (cf *CarFactory) CreateTransport() Transport {
	return &Car{}
}

// BikeFactory is a concrete creator that returns a Bike.
type BikeFactory struct{}

func (bf *BikeFactory) CreateTransport() Transport {
	return &Bike{}
}

//==============================
//OR

// TransportFactory is the factory method that returns a Transport
// func TransportFactory(transportType string) Transport {
// 	if transportType == "car" {
// 		return &Car{}
// 	}
// 	if transportType == "bike" {
// 		return &Bike{}
// 	}
// 	return nil
// }

//==============================
/*
Advantages of Factory Method Pattern:
1-Flexibility: You can add new types of products without modifying existing code (e.g., you can add a Truck class later).

2- Encapsulation: The creation logic is abstracted from the client, so the client only interacts with interfaces.

3- Single Responsibility Principle: The factory method takes care of creating objects, while other parts of the system can focus on using the objects.
*/

// ==============================================
func main() {
	// Create a Car using the CarFactory
	var factory TransportFactory = &CarFactory{}
	car := factory.CreateTransport()
	fmt.Println(car.Deliver())

	// Create a Bike using the BikeFactory
	factory = &BikeFactory{}
	bike := factory.CreateTransport()
	fmt.Println(bike.Deliver())
	// car := TransportFactory("car")
	// bike := TransportFactory("bike")

	// fmt.Println(car.Deliver())  // Output: Delivering by car
	// fmt.Println(bike.Deliver()) // Output: Delivering by bike
}
