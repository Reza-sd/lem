package room

import (
	"fmt"
	"math/rand"
	//"crypto/rand"
)

func (r *Room) OneRandomNextRoom() (Mtr, error) {
	//--------------------------------------
	lenConnectionSlice := len(r.getConnectionSlice())
	if lenConnectionSlice == 0 {
		return Mtr(0), fmt.Errorf("empty")
	}
	//-------------------------------------
	randomNextRoomIndex := rand.Intn(lenConnectionSlice)

	nextRandomRoomName := r.getConnectionSlice()[randomNextRoomIndex]
	//-----------------------
	return nextRandomRoomName, nil
}
