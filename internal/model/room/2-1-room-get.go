package room

import (
	"math/rand"
	//"crypto/rand"
)

// ==================================================
func (get *getter) name() mtr {
	return get.room.name
}

func (get *getter) allSeats() mtr {
	return get.room.allSeats
}
func (get *getter) usedSeats() mtr {
	return get.room.usedSeats
}
func (get *getter) connectionSlice() []mtr {
	return get.room.connectionSlice
}

//================hasOneFreeSeat===============================

func (get *getter) hasOneFreeSeat() bool {

	return get.usedSeats() < get.allSeats()
}

//==================OneRandomNextRoom=============================

func (get *getter) OneRandomNextRoom() mtr {

	lenConnectionSlice := len(get.connectionSlice())
	if lenConnectionSlice == 0 {

		return Answer[mtr](0, Room_get_OneRandomNextRoom, Room_get_OneRandomNextRoom_code_10)
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	if randomNextRoomIndex >= lenConnectionSlice {
		return Answer[mtr](0, Room_get_OneRandomNextRoom, Room_get_OneRandomNextRoom_Code_20)
	}

	nextRandomRoomName := get.connectionSlice()[randomNextRoomIndex]
	return Answer[mtr](nextRandomRoomName, Room_get_OneRandomNextRoom, null)
}

// ==========================================================
