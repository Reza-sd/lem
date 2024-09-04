package room

// ================================
type room struct {
	name            Mtr
	allSeats        Mtr
	usedSeats       Mtr
	connectionSlice []Mtr
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

// -----------------------
func newRoom() *room { //Constructor=factory function
	r := &room{}
	r.get = getter{room: r}
	r.set = setter{room: r}
	r.update = updater{room: r}
	return r
}

//------------------------------------------
func roomBuilder(name, endRoomName Mtr, connectionSlice []Mtr) *room {
	r := newRoom()
	r.set.name(name).connectionSlice(connectionSlice)
	if name == startRoomName || name == endRoomName {
		r.set.allSeats(MaxSeatsStartEnd).usedSeats(UsedSeatsStartEnd)
	} else {
		r.set.allSeats(AllSeatsNormalRoom).usedSeats(0)

	}
	return r
	// if its first then? if end then
}

// ---------------------------------------
