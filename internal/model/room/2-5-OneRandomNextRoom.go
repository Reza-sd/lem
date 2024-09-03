package room

import (
	"fmt"
	"math/rand"
	//"crypto/rand"
)

// ===============================================================
func (get *getter) OneRandomNextRoom() (Mtr, error) {
	//--------------------------------------
	lenConnectionSlice := len(get.room.get.connectionSlice())
	if lenConnectionSlice == 0 {
		return Mtr(0), fmt.Errorf("empty")
	}
	//-------------------------------------
	randomNextRoomIndex := rand.Intn(lenConnectionSlice)

	nextRandomRoomName := get.room.get.connectionSlice()[randomNextRoomIndex]
	//-----------------------
	return nextRandomRoomName, nil
}

// ===============================================================
