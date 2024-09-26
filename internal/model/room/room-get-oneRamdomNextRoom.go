package room

import (
	"math/rand"
	//"crypto/rand"
)

const (
	//description for log only
	_Get_OneRandomNextRoom    = "Get.OneRandomNextRoom."
	_Get_OneRandomNextRoom_10 = _Get_OneRandomNextRoom + "10:EmptyConnectionSlice"
)

//==============================

func (get *getT) OneRandomNextRoom() (rT, []errT) {
	lenConnectionSlice := len(get.ConnectionSlice())
	if lenConnectionSlice == 0 {
		return 0, logger.Err.Rlog(Get_OneRandomNextRoom_10, nil)
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	nextRandomRoomName := get.ConnectionSlice()[randomNextRoomIndex]
	return nextRandomRoomName, nil
}

//================================================
