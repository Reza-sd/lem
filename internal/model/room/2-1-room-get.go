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

func (r *room) GetOneRandomNextRoom() (m, []uint8) {

	lenConnectionSlice := len(r.GetConnectionSlice())
	if lenConnectionSlice == 0 {
		return 0, statusWrapper(GetOneRandomNextRoom10, nil)
		//return Answer[m](0, GetOneRandomNextRoom, GetOneRandomNextRoom10, r)
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	if randomNextRoomIndex >= lenConnectionSlice {
		return 0, statusWrapper(GetOneRandomNextRoom10, nil)
		//return Answer[m](0, GetOneRandomNextRoom, GetOneRandomNextRoom20, r)
	}

	nextRandomRoomName := r.GetConnectionSlice()[randomNextRoomIndex]
	return nextRandomRoomName, nil
	//return Answer[m](nextRandomRoomName, GetOneRandomNextRoom, Null, r)
}

// ==========================================================
