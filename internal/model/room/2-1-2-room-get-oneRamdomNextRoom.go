package room

import (
	"math/rand"
	//"crypto/rand"
)
//==================OneRandomNextRoom=============================

func (get *roomGetter) OneRandomNextRoom() (RT, statTypeArr) {

	lenConnectionSlice := len(get.ConnectionSlice())
	if lenConnectionSlice == 0 {
		return 0, wrapper(GetOneRandomNextRoom10, nil)
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	if randomNextRoomIndex >= lenConnectionSlice {
		return 0, wrapper(GetOneRandomNextRoom10, nil)
	}

	nextRandomRoomName := get.ConnectionSlice()[randomNextRoomIndex]
	return nextRandomRoomName, nil
}

// ==========================================================