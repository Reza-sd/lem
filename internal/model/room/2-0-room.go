package room

// =======================================================
type room struct {
	name            RT
	allSeats        RT
	usedSeats       RT
	connectionSlice []RT

	//statusLine []ET // to store current statusLine, to ckeck if this instance healthy or not
}

type rmBuildArg struct {
	name            RT
	endRoomName     RT
	connectionSlice []RT
}

// ------------------------------------
// func (r *room)wrapper(statLevel uint8, statCode uint8, preStatCodesSlice []uint8) []uint8
func wrapper(statCode uint8, preStatCodesSlice []uint8) []uint8 {
	if statCode == 0 && len(preStatCodesSlice) == 0 {
		return nil
	}

	if preStatCodesSlice == nil {
		preStatCodesSlice = []uint8{}
	}
	return append(preStatCodesSlice, statCode)
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
		//statusLine:      nil,
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
	//var a []int
	//print(len(a))
	// println(a==nil)
	// a=[]int{1}
	// println(a[0])
	// clear(a)
	// println(a)
	// println(a==nil)
	// a=nil
	// println(a==nil)
}

//=============================================================
