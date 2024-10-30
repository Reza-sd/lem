package main

import "fmt"

type BuilderType uint8

const (
	_NORMAL_BUILDER BuilderType = iota
	_IGLOO_BUILDER
)

//===================================
// 1- Define Product

type house struct {
	windowType string
	doorType   string
	floor      int
}

//===================================
// 2- Builder interface
type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

//===================================
// 3- Concrete Builders
//builder A
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// func newNormalBuilder() *normalBuilder {
//     return &normalBuilder{}
// }

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

//----------------------
//builder B
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// func newIglooBuilder() *iglooBuilder {
//     return &iglooBuilder{}
// }

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

//===================================
// Director (optional)
type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

//===================================
//_normalBuilder BuilderType = iota

func getBuilder(builderType BuilderType) iBuilder {
	if builderType == _NORMAL_BUILDER {
		return &normalBuilder{}
	}
	if builderType == _IGLOO_BUILDER {
		return &iglooBuilder{}
	}
	return nil
}

//===================================
func main() {
	normalBuilder := getBuilder(_NORMAL_BUILDER)
	iglooBuilder := getBuilder(_IGLOO_BUILDER)

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}

//===========================================
