package room

import (
	"fmt"
	"math/rand"
	//"crypto/rand"
)

// ==============================
func (r *room) Print() {

	fmt.Printf("\nRoom: Name=%v, AllSeats=%v, UsedSeats=%v, ConnectionSlice=%v\n", r.get.name(), r.get.allSeats(), r.get.usedSeats(), r.get.connectionSlice())

}

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

func (get *getter) OneRandomNextRoom() answer[mtr] {

	lenConnectionSlice := len(get.connectionSlice())
	if lenConnectionSlice == 0 {
		return answer[mtr]{Code: emptySlice, Msg: "empty slice"}
		//return answer[mtr]{sCode: emptySlice, sMsg: "empty slice"}
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	if randomNextRoomIndex >= lenConnectionSlice {
		return answer[mtr]{Code: sliceOverFlow, Msg: "index does not exist"}
	}

	nextRandomRoomName := get.connectionSlice()[randomNextRoomIndex]

	return answer[mtr]{Ans: nextRandomRoomName}
}

//==========================================================
