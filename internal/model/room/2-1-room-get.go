package room

import (
	"math/rand"
	//"crypto/rand"
)

// ==================================================
func (r *room) GetName() RT {
	return r.name
}

func (r *room) GetAllSeats() RT {
	return r.allSeats
}
func (r *room) GetUsedSeats() RT {
	return r.usedSeats
}
func (r *room) GetConnectionSlice() []RT {
	return r.connectionSlice
}

//================hasOneFreeSeat===============================

func (r *room) hasOneFreeSeat() bool {

	return r.GetUsedSeats() < r.GetAllSeats()
}

//==================OneRandomNextRoom=============================

func (r *room) GetOneRandomNextRoom() (RT, []Err) {

	lenConnectionSlice := len(r.GetConnectionSlice())
	if lenConnectionSlice == 0 {
		return 0, Wrapper(GetOneRandomNextRoom10, nil)
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	if randomNextRoomIndex >= lenConnectionSlice {
		return 0, Wrapper(GetOneRandomNextRoom10, nil)
	}

	nextRandomRoomName := r.GetConnectionSlice()[randomNextRoomIndex]
	return nextRandomRoomName, nil
}

// ==========================================================
