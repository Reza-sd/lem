package room

import (
	"math/rand"
	//"crypto/rand"
)

// ==================================================
func (get *roomGetter) Name() RT {
	return get.room.data.name //r.data.name
}

func (get *roomGetter) AllSeats() RT {
	return get.room.data.allSeats
}
func (get *roomGetter) UsedSeats() RT {
	return get.room.data.usedSeats
}
func (get *roomGetter) ConnectionSlice() []RT {
	return get.room.data.connectionSlice
}

// func (d *roomData)mio() {

// }
//================hasOneFreeSeat===============================

func (get *roomGetter) hasOneFreeSeat() bool {

	return get.UsedSeats() < get.AllSeats()
}

//==================OneRandomNextRoom=============================

func (get *roomGetter) GetOneRandomNextRoom() (RT, []statType) {

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
