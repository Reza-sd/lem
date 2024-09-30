package main

import "fmt"

/*
The Builder Pattern is a creational design pattern that helps construct complex objects step-by-step. Unlike the Factory Method, which deals with creating objects in one step, the Builder Pattern is useful when the creation process involves multiple steps or requires different configurations. It separates the construction of an object from its representation, allowing the same construction process to create different representations.

Example: Building a Computer
In this example, we'll build a Computer object step by step, which has multiple components like CPU, RAM, Storage, and GPU. The Builder will allow us to configure and build different kinds of computers.
*/

//==============================================
//1. Define the Product (Computer):
// Computer is the product that is constructed using the builder
type Computer struct {
	CPU     string
	RAM     string
	Storage string
	GPU     string
}

func (c *Computer) String() string {
	return fmt.Sprintf("Computer Specifications:\nCPU: %s\nRAM: %s\nStorage: %s\nGPU: %s", c.CPU, c.RAM, c.Storage, c.GPU)
}

//==============================================
//2. Define the Builder Interface:
//We will create an interface for the builder which includes the different methods that can be used to build the computer.

// ComputerBuilder defines the interface for building a computer
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetRAM(ram string) ComputerBuilder
	SetStorage(storage string) ComputerBuilder
	SetGPU(gpu string) ComputerBuilder

	Build() Computer
}

//==============================================
//3. Concrete Builder (PCBuilder):
//The PCBuilder will implement the ComputerBuilder interface, providing concrete implementations of the build steps.

// PCBuilder is the concrete builder that constructs the computer
type PCBuilder struct {
	cpu     string
	ram     string
	storage string
	gpu     string
}

// NewPCBuilder returns a new instance of PCBuilder
func NewPCBuilder() *PCBuilder {
	return &PCBuilder{}
}

func (b *PCBuilder) SetCPU(cpu string) ComputerBuilder {
	b.cpu = cpu
	return b
}

func (b *PCBuilder) SetRAM(ram string) ComputerBuilder {
	b.ram = ram
	return b
}

func (b *PCBuilder) SetStorage(storage string) ComputerBuilder {
	b.storage = storage
	return b
}

func (b *PCBuilder) SetGPU(gpu string) ComputerBuilder {
	b.gpu = gpu
	return b
}

func (b *PCBuilder) Build() Computer {
	return Computer{
		CPU:     b.cpu,
		RAM:     b.ram,
		Storage: b.storage,
		GPU:     b.gpu,
	}
}

//===========================================
/*
Advantages of the Builder Pattern:

- Separation of Concerns: The creation logic is separated from the object itself, making it easier to manage the construction of complex objects.

- Fluent API: The method chaining allows for a clean, readable API for object construction.

- Immutability: Once the Computer is built, it is immutable (since we don't expose setters), which can help prevent unintended changes.

- Scalability: You can easily add more methods to the builder to configure additional properties of the object.
*/

//==============================================

func main() {
	// Create a Gaming PC using the builder pattern
	gamingPC := NewPCBuilder().
		SetCPU("Intel Core i9").
		SetRAM("32GB").
		SetStorage("1TB SSD").
		SetGPU("NVIDIA RTX 3080").
		Build()

	fmt.Println(gamingPC)

	// Create an Office PC with different specs
	officePC := NewPCBuilder().
		SetCPU("Intel Core i5").
		SetRAM("16GB").
		SetStorage("512GB SSD").
		SetGPU("Intel Integrated Graphics").
		Build()

	fmt.Println(officePC)
}

//==============================================
