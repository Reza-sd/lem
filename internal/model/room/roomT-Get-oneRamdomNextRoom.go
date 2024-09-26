package room

import (
	"math/rand"
	//"crypto/rand"
)

//==============================

func (get *roomTgetT) OneRandomNextRoom() (rT, []errT) {
	lenConnectionSlice := len(get.ConnectionSlice())
	if lenConnectionSlice == 0 {
		return 0, logger.Err.Rlog(roomTgetT_OneRandomNextRoom_10, nil)
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	nextRandomRoomName := get.ConnectionSlice()[randomNextRoomIndex]
	return nextRandomRoomName, nil
}

//================================================
