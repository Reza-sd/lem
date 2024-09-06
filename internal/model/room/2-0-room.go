package room

import (
	"fmt"
)

// =======================================================
type room struct {
	name            m
	allSeats        m
	usedSeats       m
	connectionSlice []m

	Errdb map[e]e
}

type rmBuildArg struct {
	name            m
	endRoomName     m
	connectionSlice []m
}

// -------------------------------------------------------------
func newPlainRoom() *room { //Constructor=factory function=builder
	return &room{
		connectionSlice: []m{},
		Errdb:           map[e]e{Null: Null},
	}
}

// -----------------------------------------------------------
func newRuledRoom(rm rmBuildArg) *room { //Constructor=factory function=builder
	r := newPlainRoom().SetName(rm.name).SetConnectionSlice(rm.connectionSlice)
	if rm.name == startRoomName || rm.name == rm.endRoomName {
		r.SetAllSeats(MaxSeatsStartEnd).SetUsedSeats(UsedSeatsStartEnd)
	} else {
		r.SetAllSeats(AllSeatsNormalRoom).SetUsedSeats(0)

	}

	return r
	// if its first then? if end then
}

// -------------------------------
func Answer[T any](returnedValue T, funcNameCode, statusCode e, r *room) T {
	r.Errdb[funcNameCode] = statusCode
	return returnedValue
}

// ---------------------------------------------------------------
func (r *room) Print() {

	fmt.Printf("\nRoom: Name=%v, AllSeats=%v, UsedSeats=%v, ConnectionSlice=%v\n", r.GetName(), r.GetAllSeats(), r.GetUsedSeats(), r.GetConnectionSlice())

}

//=============================================================
