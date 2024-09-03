package room

import (
	"fmt"
	"math/rand"
	//"crypto/rand"
)

// ===============================================================
func (r *Room) OneRandomNextRoom() (Mtr, error) {
	//--------------------------------------
	lenConnectionSlice := len(r.get.connectionSlice())
	if lenConnectionSlice == 0 {
		return Mtr(0), fmt.Errorf("empty")
	}
	//-------------------------------------
	randomNextRoomIndex := rand.Intn(lenConnectionSlice)

	nextRandomRoomName := r.get.connectionSlice()[randomNextRoomIndex]
	//-----------------------
	return nextRandomRoomName, nil
}

// ===============================================================
