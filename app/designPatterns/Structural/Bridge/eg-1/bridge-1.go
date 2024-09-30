package main

import "fmt"

//===========================================
/*
The Bridge Pattern is a structural design pattern that helps to decouple an abstraction from its implementation, allowing both to vary independently. This pattern is particularly useful when you have an abstraction that can have multiple implementations, and you want to be able to switch between them at runtime.
*/

//==========================================
// Implementor interface
type Device interface {
	On()
	Off()
	SetVolume(volume int)
}

// Concrete Implementor 1
type TV struct {
	volume int
}

func (tv *TV) On() {
	fmt.Println("TV is On")
}

func (tv *TV) Off() {
	fmt.Println("TV is Off")
}

func (tv *TV) SetVolume(volume int) {
	tv.volume = volume
	fmt.Printf("TV volume set to %d\n", tv.volume)
}

//==========================================
// Concrete Implementor 2
type Radio struct {
	volume int
}

func (r *Radio) On() {
	fmt.Println("Radio is On")
}

func (r *Radio) Off() {
	fmt.Println("Radio is Off")
}

func (r *Radio) SetVolume(volume int) {
	r.volume = volume
	fmt.Printf("Radio volume set to %d\n", r.volume)
}

//==========================================
// Abstraction
type RemoteControl struct {
	device Device
}

func (rc *RemoteControl) TurnOn() {
	rc.device.On()
}

func (rc *RemoteControl) TurnOff() {
	rc.device.Off()
}

func (rc *RemoteControl) VolumeUp() {
	rc.device.SetVolume(10)
}

func (rc *RemoteControl) VolumeDown() {
	rc.device.SetVolume(0)
}

//======================
/*
In this example:

Device is the implementor interface.
TV and Radio are concrete implementors.
RemoteControl is the abstraction that uses a Device.
*/
//==========================================
func main() {
	tv := &TV{}
	radio := &Radio{}

	remote := &RemoteControl{device: tv}
	remote.TurnOn()
	remote.VolumeUp()

	remote.device = radio
	remote.TurnOn()
	remote.VolumeDown()
}

//==========================================
