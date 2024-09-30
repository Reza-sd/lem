package main

import "fmt"

// Product interface
type Vehicle interface {
	Drive()
}

// Concrete Products
type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Driving a car")
}

type Bike struct{}

func (b *Bike) Drive() {
	fmt.Println("Riding a bike")
}

// Creator interface
type VehicleFactory interface {
	CreateVehicle() Vehicle
}

// Concrete Creators
type CarFactory struct{}

func (cf *CarFactory) CreateVehicle() Vehicle {
	return &Car{}
}

type BikeFactory struct{}

func (bf *BikeFactory) CreateVehicle() Vehicle {
	return &Bike{}
}

// Client code
func main() {
	carFactory := &CarFactory{}
	bikeFactory := &BikeFactory{}

	car := carFactory.CreateVehicle()
	bike := bikeFactory.CreateVehicle()

	car.Drive()
	bike.Drive()
}
