package room

// =======================================================
type room struct {
	name            RT
	allSeats        RT
	usedSeats       RT
	connectionSlice []RT

	statusLine []ET
}

type rmBuildArg struct {
	name            RT
	endRoomName     RT
	connectionSlice []RT
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
		connectionSlice: []RT{},
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

	println()

	print("Room: Name=")
	print(r.GetName())

	print(", AllSeats=")
	print(r.GetAllSeats())

	print(", UsedSeats=")
	print(r.GetUsedSeats())

	print(", ConnectionSlice[]index:(value)=")
	for index, value := range r.GetConnectionSlice() {
		print(index, ":(", value, "),")

	}

	println()

}

//=============================================================
