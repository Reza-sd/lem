package room

import (
	"math/rand"
	//"crypto/rand"
)

// ==================================================
func (g *getter) name() mtr {
	return g.room.name
}

func (g *getter) allSeats() mtr {
	return g.room.allSeats
}
func (g *getter) usedSeats() mtr {
	return g.room.usedSeats
}
func (g *getter) connectionSlice() []mtr {
	return g.room.connectionSlice
}

//================hasOneFreeSeat===============================

func (g *getter) hasOneFreeSeat() bool {

	return g.usedSeats() < g.allSeats()
}

//==================OneRandomNextRoom=============================

func (g *getter) OneRandomNextRoom() mtr {

	lenConnectionSlice := len(g.connectionSlice())
	if lenConnectionSlice == 0 {

		return Answer[mtr](0, Room_get_OneRandomNextRoom, Room_get_OneRandomNextRoom_code_10, g.room)
	}

	randomNextRoomIndex := rand.Intn(lenConnectionSlice) // len 4 => random :0,1,2,3

	if randomNextRoomIndex >= lenConnectionSlice {
		return Answer[mtr](0, Room_get_OneRandomNextRoom, Room_get_OneRandomNextRoom_Code_20, g.room)
	}

	nextRandomRoomName := g.connectionSlice()[randomNextRoomIndex]
	return Answer[mtr](nextRandomRoomName, Room_get_OneRandomNextRoom, null, g.room)
}

// ==========================================================
