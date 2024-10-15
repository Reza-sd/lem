package room

import (
	"math/rand"
	//"crypto/rand"
)

// ==============================
//const _EMPTY_CONNECTION_SLICE = 0

func (get *roomTgetT) OneRandomNextRoom() (rT, []errT) {
	lenConnectionSlice := len(get.ConnectionSlice())
	if lenConnectionSlice == _EMPTY_CONNECTION_SLICE {
		return 0, logger.Act.Err.Rlog(_ERR_roomT_getT_OneRandomNextRoom_10, nil, "if lenConnectionSlice == 0")
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	nextRandomRoomName := get.ConnectionSlice()[randomNextRoomIndex]
	return nextRandomRoomName, nil
}

//================================================
