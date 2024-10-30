package main

import "fmt"


//============Global========================
type builderType_uint8 uint8

const (
	_NORMAL_BUILDER builderType_uint8 = iota
	_IGLOO_BUILDER
)

var (
	windowType_Map_builderType_uint8_string = map[builderType_uint8]string{
		_NORMAL_BUILDER: "Wooden Window",
		_IGLOO_BUILDER:  "snow window",
	}

	doorType_Map_builderType_uint8_string = map[builderType_uint8]string{
		_NORMAL_BUILDER: "Wooden door",
		_IGLOO_BUILDER:  "Snow Door",
	}
)

//===================================
// 1- Define Product

type house_struct struct {
	windowType_string string
	doorType_string   string
	floor_uint8      uint8
}

//===================================
// 2- Builder interface
type Builder_interface interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	buildHouse() *house_struct
}

//===================================
// 3- Concrete Builders
//builder A
type normalBuilder_struct struct {
	windowType_string string
	doorType_string   string
	floor_uint8      uint8
}

func newNormalBuilder() *normalBuilder_struct {
	return new(normalBuilder_struct)
}

func (b *normalBuilder_struct) setWindowType() {
	b.windowType_string = windowType_Map_builderType_uint8_string[_NORMAL_BUILDER]
}

func (b *normalBuilder_struct) setDoorType() {
	b.doorType_string = doorType_Map_builderType_uint8_string[_NORMAL_BUILDER]
}

func (b *normalBuilder_struct) setNumFloor() {
	b.floor_uint8 =uint8(2)
}

func (b *normalBuilder_struct) buildHouse() *house_struct {
	h := new(house_struct)
	h.doorType_string = b.doorType_string
	h.windowType_string = b.windowType_string
	h.floor_uint8 = b.floor_uint8
	return h
}

//----------------------
//builder B
type iglooBuilder_struct struct {
	windowType_string string
	doorType_string   string
	floor_uint8      uint8
}

func newIglooBuilder() *iglooBuilder_struct {
	return new(iglooBuilder_struct)
}

func (b *iglooBuilder_struct) setWindowType() {
	b.windowType_string = windowType_Map_builderType_uint8_string[_IGLOO_BUILDER]
}

func (b *iglooBuilder_struct) setDoorType() {
	b.doorType_string = doorType_Map_builderType_uint8_string[_IGLOO_BUILDER]
}

func (b *iglooBuilder_struct) setNumFloor() {
	b.floor_uint8 = uint8(1)
}

func (b *iglooBuilder_struct) buildHouse() *house_struct {
	h := new(house_struct)
	h.doorType_string = b.doorType_string
	h.windowType_string = b.windowType_string
	h.floor_uint8 = b.floor_uint8

	return h
}

//===================================
// Director (optional)
type director_struct struct {
	builder Builder_interface
}

func newDirector(b Builder_interface) *director_struct {
	d := new(director_struct)
	d.builder = b
	return d
}

func (d *director_struct) setBuilder(b Builder_interface) {
	d.builder = b
}

func (d *director_struct) buildHouse() *house_struct {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.buildHouse()
}

//===================================

func getBuilder(builderType builderType_uint8) Builder_interface {
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

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType_string)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType_string)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor_uint8)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType_string)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType_string)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor_uint8)
}

//===========================================
