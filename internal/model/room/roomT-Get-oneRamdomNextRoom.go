package room

import (
	"math/rand"
	//"crypto/rand"
)

//==============================

func (get *roomTgetT) OneRandomNextRoom() (rT, []errT) {
	lenConnectionSlice := len(get.ConnectionSlice())
	if lenConnectionSlice == 0 {
		return 0, logger.Act.Err.Rlog(_roomTgetT_OneRandomNextRoom_10, nil, "if lenConnectionSlice == 0")
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	nextRandomRoomName := get.ConnectionSlice()[randomNextRoomIndex]
	return nextRandomRoomName, nil
}

//================================================
