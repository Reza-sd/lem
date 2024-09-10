package room

// =======================================================
type room struct {
	data Data

	get roomGetter
	set roomSetter
	act roomAction
}
type Data struct {
	name            RT
	allSeats        RT
	usedSeats       RT
	connectionSlice []RT
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

// ------------------------------------
type rmBuildArg struct {
	name            RT
	endRoomName     RT
	connectionSlice []RT
}

// ------------------------------------
func wrapper(statCode statCodeType, preStatCodesSlice statArrTyp) statArrTyp {
	if statCode == 0 && len(preStatCodesSlice) == 0 {
		return nil
	}

	if preStatCodesSlice == nil {
		preStatCodesSlice = statArrTyp{}
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
func newRuledRoom(rm rmBuildArg) (*room, statArrTyp) { //Constructor=factory function=builder
	r := newPlainRoom()

	if err := r.set.Name(rm.name); err != nil {
		return nil, wrapper(100, err)
	}
	r.set.ConnectionSlice(rm.connectionSlice)

	if rm.name == startRoomName || rm.name == rm.endRoomName {

		r.set.AllSeats(MaxSeatsStartEnd)
		r.set.UsedSeats(UsedSeatsStartEnd)
	} else {
		r.set.AllSeats(AllSeatsNormalRoom)
		r.set.UsedSeats(0)

	}

	return r, nil
	// if its first then? if end then
}

//
