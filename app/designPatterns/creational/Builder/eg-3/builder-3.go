package main

import "fmt"
//---------------------------------------
// Product
type House struct {
	walls   int
	doors   int
	windows int
	roof    string
	garage  bool
}
//---------------------------------------
// Builder interface
type HouseBuilder interface {
	SetWalls(walls int) HouseBuilder
	SetDoors(doors int) HouseBuilder
	SetWindows(windows int) HouseBuilder
	SetRoof(roof string) HouseBuilder
	SetGarage(garage bool) HouseBuilder

	Build() *House
}
//---------------------------------------
// Concrete Builder
type ConcreteHouseBuilder struct {
	house *House
}

func NewConcreteHouseBuilder() *ConcreteHouseBuilder {
	return &ConcreteHouseBuilder{house: &House{}}
}

func (b *ConcreteHouseBuilder) SetWalls(walls int) HouseBuilder {
	b.house.walls = walls
	return b
}

func (b *ConcreteHouseBuilder) SetDoors(doors int) HouseBuilder {
	b.house.doors = doors
	return b
}

func (b *ConcreteHouseBuilder) SetWindows(windows int) HouseBuilder {
	b.house.windows = windows
	return b
}

func (b *ConcreteHouseBuilder) SetRoof(roof string) HouseBuilder {
	b.house.roof = roof
	return b
}

func (b *ConcreteHouseBuilder) SetGarage(garage bool) HouseBuilder {
	b.house.garage = garage
	return b
}

func (b *ConcreteHouseBuilder) Build() *House {
	return b.house
}
//---------------------------------------
// Director (optional)
type HouseDirector struct {
	builder HouseBuilder
}

func NewHouseDirector(b HouseBuilder) *HouseDirector {
	return &HouseDirector{builder: b}
}

func (d *HouseDirector) ConstructSimpleHouse() *House {
	return d.builder.SetWalls(4).SetDoors(1).SetWindows(4).SetRoof("Tiles").Build()
}

func (d *HouseDirector) ConstructLuxuryHouse() *House {
	return d.builder.SetWalls(6).SetDoors(3).SetWindows(8).SetRoof("Slate").SetGarage(true).Build()
}
//---------------------------------------
func main() {
	builder := NewConcreteHouseBuilder()

	// Using builder directly
	house1 := builder.SetWalls(4).SetDoors(2).SetWindows(6).SetRoof("Shingles").Build()
	fmt.Printf("House 1: %+v\n", house1)

	// Using director
	director := NewHouseDirector(builder)

	simpleHouse := director.ConstructSimpleHouse()
	fmt.Printf("Simple House: %+v\n", simpleHouse)

	luxuryHouse := director.ConstructLuxuryHouse()
	fmt.Printf("Luxury House: %+v\n", luxuryHouse)
}
//---------------------------------------
