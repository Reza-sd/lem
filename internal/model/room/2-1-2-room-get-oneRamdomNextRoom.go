package room

import (
	"math/rand"
	//"crypto/rand"
)

//==================OneRandomNextRoom=============================

func (get *rGetter) OneRandomNextRoom() (RT, statArrT) {

	lenConnectionSlice := len(get.ConnectionSlice())
	if lenConnectionSlice == 0 {
		return 0, stat(Get_OneRandomNextRoom_EmptyConnectionSlice, nil)
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	nextRandomRoomName := get.ConnectionSlice()[randomNextRoomIndex]
	return nextRandomRoomName, nil
}

// ==========================================================
