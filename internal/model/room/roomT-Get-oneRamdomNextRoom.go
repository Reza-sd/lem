package room

import (
	"math/rand"
	//"crypto/rand"
)

const (
	//description for log only
	_roomT_getT_OneRandomNextRoom    = "Get.OneRandomNextRoom."
	_roomT_getT_OneRandomNextRoom_10 = _roomT_getT_OneRandomNextRoom + "10:EmptyConnectionSlice"
)

//==============================

func (get *roomTgetT) OneRandomNextRoom() (rT, []errT) {
	lenConnectionSlice := len(get.ConnectionSlice())
	if lenConnectionSlice == 0 {
		return 0, logger.Err.Rlog(roomT_getT_OneRandomNextRoom_10, nil)
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	nextRandomRoomName := get.ConnectionSlice()[randomNextRoomIndex]
	return nextRandomRoomName, nil
}

//================================================
