package room

// =======================================================
type room struct {
	data Data //private

	//categorise behaviour
	set rSetter //private

	Get rGetter //export (public)
	Act rAction //export (public)
}
type Data struct {
	name            RT
	allSeats        RT
	usedSeats       RT
	connectionSlice []RT
}

type rGetter struct {
	room *room
}
type rSetter struct {
	room *room
}
type rAction struct {
	room *room
}

//==========================================================

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

	r.Get.room = r
	r.set.room = r
	r.Act.room = r

	return r
}

// ------------------------------------
type rmBuildArg struct {
	name            RT
	endRoomName     RT
	connectionSlice []RT
}

func newRuledRoom(rm rmBuildArg) (*room, statArrT) { //Constructor=factory function=builder
	r := newPlainRoom()

	if err := r.set.Name(rm.name); err != nil {
		return nil, stat(1, err)
	}
	if err := r.set.ConnectionSlice(rm.connectionSlice); err != nil {
		return nil, stat(2, err)
	}

	if rm.name == startRoomName || rm.name == rm.endRoomName {

		if err := r.set.AllSeats(MaxSeatsStartEnd); err != nil {
			return nil, stat(3, err)

		}
		if err := r.set.UsedSeats(UsedSeatsStartEnd); err != nil {
			return nil, stat(4, err)

		}
	} else {
		if err := r.set.AllSeats(AllSeatsNormalRoom); err != nil {
			return nil, stat(5, err)

		}
		if err := r.set.UsedSeats(0); err != nil {
			return nil, stat(6, err)

		}

	}

	return r, nil
	// if its first then? if end then
}

//---------------------------------------
