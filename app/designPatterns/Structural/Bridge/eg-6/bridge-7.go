package main

import "fmt"

//======================================
/*
The Bridge Pattern is a structural design pattern used to decouple an abstraction from its implementation, allowing the two to evolve independently. In simpler terms, it helps in separating the interface from its concrete implementation, enabling flexibility and reusability.

# Key Concepts in Bridge Pattern:

1- Abstraction: The interface or abstract class that defines the higher-level control logic.

2- Implementation: Concrete implementations that the abstraction can use without being tightly coupled to them.

# When to Use the Bridge Pattern:

- When you have multiple variants of an object and you don't want them tightly coupled.
- When you want to change implementation details without affecting the client code.
- When the abstraction and implementation need to evolve independently.

# Bridge Pattern in Golang:

Golang doesn't have traditional inheritance, but the Bridge pattern can be implemented using interfaces and struct composition.
*/
//====================================
//1. Define the Implementor Interface
// This interface will define the methods that concrete implementations must have.

// Implementor interface
type Device interface {
	TurnOn()
	TurnOff()
	SetVolume(volume int)
}

//====================================
//2. Create Concrete Implementations of the Device
//These will be the concrete implementations of the Device interface.

// Concrete Implementor 1: TV
type TV struct {
	isOn   bool
	volume int
}

func (t *TV) TurnOn() {
	t.isOn = true
	fmt.Println("TV is now ON")
}

func (t *TV) TurnOff() {
	t.isOn = false
	fmt.Println("TV is now OFF")
}

func (t *TV) SetVolume(volume int) {
	t.volume = volume
	fmt.Printf("TV volume set to %d\n", volume)
}

// Concrete Implementor 2: Radio
type Radio struct {
	isOn   bool
	volume int
}

func (r *Radio) TurnOn() {
	r.isOn = true
	fmt.Println("Radio is now ON")
}

func (r *Radio) TurnOff() {
	r.isOn = false
	fmt.Println("Radio is now OFF")
}

func (r *Radio) SetVolume(volume int) {
	r.volume = volume
	fmt.Printf("Radio volume set to %d\n", volume)
}

//==============================================================
//3. Define the Abstraction
//The abstraction defines the control logic that interacts with the Device interface.

// Abstraction: RemoteControl
type RemoteControl struct {
	device Device
}

func (r *RemoteControl) TurnOn() {
	r.device.TurnOn()
}

func (r *RemoteControl) TurnOff() {
	r.device.TurnOff()
}

func (r *RemoteControl) SetVolume(volume int) {
	r.device.SetVolume(volume)
}

//=============================================================
//4. Define a Refined Abstraction
//The refined abstraction is a subclass of the abstraction that can provide additional functionality.

// Refined Abstraction: AdvancedRemoteControl
type AdvancedRemoteControl struct {
	RemoteControl
}

func (a *AdvancedRemoteControl) Mute() {
	fmt.Println("Muting device")
	a.device.SetVolume(0)
}

// ===================================================
func main() {
	// Using TV with RemoteControl
	tv := &TV{}
	remote := &RemoteControl{device: tv}

	remote.TurnOn()
	remote.SetVolume(10)
	remote.TurnOff()

	// Using Radio with AdvancedRemoteControl
	radio := &Radio{}
	advancedRemote := &AdvancedRemoteControl{RemoteControl{device: radio}}

	advancedRemote.TurnOn()
	advancedRemote.SetVolume(5)
	advancedRemote.Mute()
	advancedRemote.TurnOff()
}

//====================================================
