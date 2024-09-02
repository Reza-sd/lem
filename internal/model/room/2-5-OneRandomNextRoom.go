package room

import (
	"fmt"
	"math/rand"
	//"crypto/rand"

)

func (myRoom *Room) OneRandomNextRoom() (Mtr, error) {
	//--------------------------------------
	lenConnectionSlice := len(myRoom.ConnectionSlice)
	if lenConnectionSlice == 0 {
		return Mtr(0), fmt.Errorf("empty")
	}
	//----------------------------------------
	var nextRandomRoomName Mtr
	//-------------------------------------
	randomNextRoomIndex := rand.Intn(lenConnectionSlice)
	nextRandomRoomName = myRoom.ConnectionSlice[randomNextRoomIndex]
	//-----------------------
	return nextRandomRoomName, nil
}
