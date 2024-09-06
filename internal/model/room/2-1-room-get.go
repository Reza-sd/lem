package room

import (
	"math/rand"
	//"crypto/rand"
)

// ==================================================
func (r *room) GetName() m {
	return r.name
}

func (r *room) GetAllSeats() m {
	return r.allSeats
}
func (r *room) GetUsedSeats() m {
	return r.usedSeats
}
func (r *room) GetConnectionSlice() []m {
	return r.connectionSlice
}

//================hasOneFreeSeat===============================

func (r *room) hasOneFreeSeat() bool {

	return r.GetUsedSeats() < r.GetAllSeats()
}

//==================OneRandomNextRoom=============================

func (r *room) GetOneRandomNextRoom() m {

	lenConnectionSlice := len(r.GetConnectionSlice())
	if lenConnectionSlice == 0 {

		return Answer[m](0, Room_get_OneRandomNextRoom, Room_get_OneRandomNextRoom_code_10, r)
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	if randomNextRoomIndex >= lenConnectionSlice {
		return Answer[m](0, Room_get_OneRandomNextRoom, Room_get_OneRandomNextRoom_Code_20, r)
	}

	nextRandomRoomName := r.GetConnectionSlice()[randomNextRoomIndex]
	return Answer[m](nextRandomRoomName, Room_get_OneRandomNextRoom, null, r)
}

// ==========================================================
