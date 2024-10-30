package main

import "fmt"


//============Global========================
type BuilderType uint8

const (
	_NORMAL_BUILDER BuilderType = iota
	_IGLOO_BUILDER
)

var (
	windowTypeMap = map[BuilderType]string{
		_NORMAL_BUILDER: "Wooden Window",
		_IGLOO_BUILDER:  "snow window",
	}

	doorTypeMap = map[BuilderType]string{
		_NORMAL_BUILDER: "Wooden door",
		_IGLOO_BUILDER:  "Snow Door",
	}
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
	buildHouse() *house
}

//===================================
// 3- Concrete Builders
//builder A
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return new(normalBuilder)
}

func (b *normalBuilder) setWindowType() {
	b.windowType = windowTypeMap[_NORMAL_BUILDER]
}

func (b *normalBuilder) setDoorType() {
	b.doorType = doorTypeMap[_NORMAL_BUILDER]
}

func (b *normalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) buildHouse() *house {
	h := new(house)
	h.doorType = b.doorType
	h.windowType = b.windowType
	h.floor = b.floor
	return h
}

//----------------------
//builder B
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return new(iglooBuilder)
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = windowTypeMap[_IGLOO_BUILDER]
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = doorTypeMap[_IGLOO_BUILDER]
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) buildHouse() *house {
	h := new(house)
	h.doorType = b.doorType
	h.windowType = b.windowType
	h.floor = b.floor

	return h
}

//===================================
// Director (optional)
type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	d := new(director)
	d.builder = b
	return d
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() *house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.buildHouse()
}

//===================================

func getBuilder(builderType BuilderType) iBuilder {
	if builderType == _NORMAL_BUILDER {
		return newNormalBuilder()
	}
	if builderType == _IGLOO_BUILDER {
		return newIglooBuilder()
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
