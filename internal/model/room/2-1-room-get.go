package room

import (
	"fmt"
	"math/rand"
	//"crypto/rand"
)

// ================================
type getter struct {
	room *room
}

// ===============================================
func (get *getter) name() Mtr {
	return get.room.name
}

func (get *getter) allSeats() Mtr {
	return get.room.allSeats
}
func (get *getter) usedSeats() Mtr {
	return get.room.usedSeats
}
func (get *getter) connectionSlice() []Mtr {
	return get.room.connectionSlice
}

//===============================================

func (get *getter) hasOneFreeSeat() bool {

	return get.usedSeats() < get.allSeats()
}

//===============================================

func (get *getter) OneRandomNextRoom() answer[Mtr] {
	//--------------------------------------
	lenConnectionSlice := len(get.room.get.connectionSlice())
	if lenConnectionSlice == 0 {
		return answer[Mtr]{Mtr(0), fmt.Errorf("empty")}
	}
	//-------------------------------------
	randomNextRoomIndex := rand.Intn(lenConnectionSlice)

	nextRandomRoomName := get.room.get.connectionSlice()[randomNextRoomIndex]
	//-----------------------
	return answer[Mtr]{nextRandomRoomName, nil}
}

//===============================================
