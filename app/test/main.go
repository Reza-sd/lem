package main

import "fmt"

// Define an interface
type getter interface {
    AllSeats() string
	UsedSeats()string
}

type setter interface {
    AllSeats() string
	UsedSeats()string
}

// Define a struct implementing the interface
type RoomData struct {
    Name string
	allSeats int
	usedSeats int
}

// Implement the Speak method for Cat
func (c RoomData) AllSeats() string {
    return "111!"
}
func (c RoomData) UsedSeats() string {
    return "222"
}

type Room struct {
    get getter
	set setter
}

func main() {
    // Create instances of Cat and Dog
    myRoomData := RoomData{Name: "Whiskers"}

    myroom := Room{get: myRoomData,set: myRoomData}
	
    fmt.Println(myroom.get.AllSeats()) 
	myroom.get.AllSeats()
	myroom.get.UsedSeats()
	myroom.set.AllSeats()
	fmt.Println(myroom)

}
