package room

// =======================================================
type room struct {
	data struct {
		name            RT
		allSeats        RT
		usedSeats       RT
		connectionSlice []RT
	}

	get roomGetter
	set roomSetter
	act roomAction
}

type roomGetter struct {
	room *room
}
type roomSetter struct {
	room *room
}
type roomAction struct {
	room *room
}

type rmBuildArg struct {
	name            RT
	endRoomName     RT
	connectionSlice []RT
}

// ------------------------------------
// func (r *room)wrapper(statLevel uint8, statCode uint8, preStatCodesSlice []uint8) []uint8
func wrapper(statCode statType, preStatCodesSlice []statType) []statType {
	if statCode == 0 && len(preStatCodesSlice) == 0 {
		return nil
	}

	if preStatCodesSlice == nil {
		preStatCodesSlice = []statType{}
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

	r := &room{}

	r.data.connectionSlice = []RT{}

	r.get.room = r
	r.set.room = r
	r.act.room = r

	return r
}

// -----------------------------------------------------------
func newRuledRoom(rm rmBuildArg) *room { //Constructor=factory function=builder
	//r := newPlainRoom().SetName(rm.name).SetConnectionSlice(rm.connectionSlice)
	r := newPlainRoom().set.Name(rm.name).ConnectionSlice(rm.connectionSlice).room
	if rm.name == startRoomName || rm.name == rm.endRoomName {

		r.set.AllSeats(MaxSeatsStartEnd).UsedSeats(UsedSeatsStartEnd)
	} else {
		r.set.AllSeats(AllSeatsNormalRoom).UsedSeats(0)

	}

	return r
	// if its first then? if end then
}

// ---------------------------------------------------------------
func (act *roomAction) Print() {

	println()

	print("Room: Name=")
	print(act.room.get.Name())

	print(", AllSeats=")
	print(act.room.get.AllSeats())

	print(", UsedSeats=")
	print(act.room.get.UsedSeats())

	print(", ConnectionSlice[]index:(value)=")
	for index, value := range act.room.get.ConnectionSlice() {
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
