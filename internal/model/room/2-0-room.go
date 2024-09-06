package room

import (
	"fmt"
)

// =======================================================
type room struct {
	name            mtr
	allSeats        mtr
	usedSeats       mtr
	connectionSlice []mtr
	//IsAvailable bool
	get getter
	set setter
	//update updater
	Errdb map[uint8]uint8
}
type getter struct {
	room *room
}

type setter struct {
	room *room
}

type rmBuildArg struct {
	name            mtr
	endRoomName     mtr
	connectionSlice []mtr
}

// -------------------------------------------------------------
func newPlainRoom() *room { //Constructor=factory function=builder
	r := &room{}
	r.get = getter{room: r}
	r.set = setter{room: r}
	r.Errdb = make(map[uint8]uint8)
	return r
}

// -----------------------------------------------------------
func newRuledRoom(rm rmBuildArg) *room { //Constructor=factory function=builder

	r := newPlainRoom()
	r.set.name(rm.name).connectionSlice(rm.connectionSlice)
	if rm.name == startRoomName || rm.name == rm.endRoomName {
		r.set.allSeats(MaxSeatsStartEnd).usedSeats(UsedSeatsStartEnd)
	} else {
		r.set.allSeats(AllSeatsNormalRoom).usedSeats(0)

	}
	return r
	// if its first then? if end then
}

// -------------------------------
func Answer[T any](returnedValue T, funcNameCode, statusCode uint8, r *room) T {
	r.Errdb[funcNameCode] = statusCode
	return returnedValue
}

// ---------------------------------------------------------------
func (r *room) Print() {

	fmt.Printf("\nRoom: Name=%v, AllSeats=%v, UsedSeats=%v, ConnectionSlice=%v\n", r.get.name(), r.get.allSeats(), r.get.usedSeats(), r.get.connectionSlice())

}

//=============================================================
