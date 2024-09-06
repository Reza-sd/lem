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

	//Errdb      map[e]e
	statusLine []uint8
}

type rmBuildArg struct {
	name            m
	endRoomName     m
	connectionSlice []m
}

// var codedb =[]uint8{0,1}
func statusWrapper(statusCode uint8, previousStatusCodeSlice []uint8) []uint8 {
	if statusCode == 0 && previousStatusCodeSlice == nil {
		return nil
	}

	if previousStatusCodeSlice == nil {
		previousStatusCodeSlice = []uint8{}
	}
	return append(previousStatusCodeSlice, statusCode)
}

// -------------------------------
// func Answer[T any](returnedValue T, funcNameCode, statusCode e, r *room) T {
// 	//statusWrapper(statusCode,r.statusLine)
// 	r.Errdb[funcNameCode] = statusCode
// 	return returnedValue
// }

// -------------------------------------------------------------
func newPlainRoom() *room { //Constructor=factory function=builder
	return &room{
		connectionSlice: []m{},
		//Errdb:           map[e]e{Null: Null},
		statusLine: nil,
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

// ---------------------------------------------------------------
func (r *room) Print() {

	fmt.Printf("\nRoom: Name=%v, AllSeats=%v, UsedSeats=%v, ConnectionSlice=%v\n", r.GetName(), r.GetAllSeats(), r.GetUsedSeats(), r.GetConnectionSlice())

}

//=============================================================
