package room

import (
	"fmt"
	"math/rand"
	//"crypto/rand"
)

func (myRoom *Room) OneRandomNextRoom() (Mtr, error) {
	//--------------------------------------
	lenConnectionSlice := len(myRoom.connectionSlice)
	if lenConnectionSlice == 0 {
		return Mtr(0), fmt.Errorf("empty")
	}
	//----------------------------------------
	var nextRandomRoomName Mtr
	//-------------------------------------
	randomNextRoomIndex := rand.Intn(lenConnectionSlice)
	nextRandomRoomName = myRoom.connectionSlice[randomNextRoomIndex]
	//-----------------------
	return nextRandomRoomName, nil
}
