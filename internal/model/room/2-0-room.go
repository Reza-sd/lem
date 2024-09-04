package room

// ================================
type room struct {
	name            mtr
	allSeats        mtr
	usedSeats       mtr
	connectionSlice []mtr
	//IsAvailable bool
	get    getter
	set    setter
	update updater
}
type getter struct {
	room *room
}

type setter struct {
	room *room
}
type updater struct {
	room *room
}
type rmBuildArg struct {
	name            mtr
	endRoomName     mtr
	connectionSlice []mtr
}

// -----------------------
func newPlainRoom() *room { //Constructor=factory function
	r := &room{}
	r.get = getter{room: r}
	r.set = setter{room: r}
	r.update = updater{room: r}
	return r
}

// ---------------------------------------
func newRuledRoom(rm rmBuildArg) *room {

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

//------------------------------------------
