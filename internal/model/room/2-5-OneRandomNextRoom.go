package room

import (
	"fmt"
	"math/rand"
)

func (myRoom *Room) OneRandomNextRoom() (Mtr, error) {
	//--------------------------------------
	lenConnectionSlice := len(myRoom.ConnectionSlice)
	if lenConnectionSlice == 0 {
		return 0, fmt.Errorf("empty")
	}
	//----------------------------------------
	var nextRandomRoomName Mtr
	//-------------------------------------
	randomNextRoomIndex := rand.Intn(lenConnectionSlice)
	nextRandomRoomName = myRoom.ConnectionSlice[randomNextRoomIndex]
	//-----------------------
	return nextRandomRoomName, nil
}
