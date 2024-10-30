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
	window_string string
	door_string   string
	floor_uint8      uint8
}

//===================================
// 2- Builder interface
type Builder_interface interface {
	set_WindowType()
	set_DoorType()
	set_NumFloor()

	get_House() *house_struct
}

//===================================
// 3- Concrete Builders
//builder A
type normalBuilder_struct struct {
	window_string string
	door_string   string
	floor_uint8      uint8
}

func newNormalBuilder() *normalBuilder_struct {
	return new(normalBuilder_struct)
}

func (b *normalBuilder_struct) set_WindowType() {
	b.window_string = windowType_Map_builderType_uint8_string[_NORMAL_BUILDER]
}

func (b *normalBuilder_struct) set_DoorType() {
	b.door_string = doorType_Map_builderType_uint8_string[_NORMAL_BUILDER]
}

func (b *normalBuilder_struct) set_NumFloor() {
	b.floor_uint8 =uint8(2)
}

func (b *normalBuilder_struct) get_House() *house_struct {
	h := new(house_struct)
	h.door_string = b.door_string
	h.window_string = b.window_string
	h.floor_uint8 = b.floor_uint8
	return h
}

//----------------------
//builder B
type iglooBuilder_struct struct {
	window_string string
	door_string   string
	floor_uint8      uint8
}

func newIglooBuilder() *iglooBuilder_struct {
	return new(iglooBuilder_struct)
}

func (b *iglooBuilder_struct) set_WindowType() {
	b.window_string = windowType_Map_builderType_uint8_string[_IGLOO_BUILDER]
}

func (b *iglooBuilder_struct) set_DoorType() {
	b.door_string = doorType_Map_builderType_uint8_string[_IGLOO_BUILDER]
}

func (b *iglooBuilder_struct) set_NumFloor() {
	b.floor_uint8 = uint8(1)
}

func (b *iglooBuilder_struct) get_House() *house_struct {
	h := new(house_struct)
	h.door_string = b.door_string
	h.window_string = b.window_string
	h.floor_uint8 = b.floor_uint8

	return h
}

//===================================
// Director (optional)
type director_struct struct {
	builder Builder_interface
}

func new_Director(b Builder_interface) *director_struct {
	d := new(director_struct)
	d.builder = b
	return d
}

func (d *director_struct) set_Builder(b Builder_interface) {
	d.builder = b
}

func (this_director *director_struct) buildHouse() *house_struct {
	this_director.builder.set_DoorType()
	this_director.builder.set_WindowType()
	this_director.builder.set_NumFloor()
	return this_director.builder.get_House()
}

//===================================

func getBuilder(this_builderType builderType_uint8) Builder_interface {
	if this_builderType == _NORMAL_BUILDER {
		return newNormalBuilder()
	}
	if this_builderType == _IGLOO_BUILDER {
		return newIglooBuilder()
	}
	return nil
}

//===================================
func main() {
	normalBuilder := getBuilder(_NORMAL_BUILDER)
	iglooBuilder := getBuilder(_IGLOO_BUILDER)

	director := new_Director(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.door_string)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.window_string)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor_uint8)

	director.set_Builder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.door_string)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.window_string)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor_uint8)
}

//===========================================
