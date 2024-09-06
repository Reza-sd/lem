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

	statusLine []e
}

type rmBuildArg struct {
	name            m
	endRoomName     m
	connectionSlice []m
}

// ------------------------------------
func Wrapper(code uint8, preCodesSlice []uint8) []uint8 {
	if code == 0 && preCodesSlice == nil {
		return nil
	}

	if preCodesSlice == nil {
		preCodesSlice = []uint8{}
	}
	return append(preCodesSlice, code)
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
		statusLine:      nil,
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
