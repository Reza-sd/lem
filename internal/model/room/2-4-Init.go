package room

// ===============================================================
func (r *Room) initiator(name, lastRoomName Mtr, connectionSlice []Mtr) {
	r.set.name(name)
	r.set.connectionSlice(connectionSlice)
	if name == 0 || name == lastRoomName {
		r.set.allSeats(MaxSeatsStartEnd)
		r.set.usedSeats(UsedSeatsStartEnd)
	} else {
		r.set.allSeats(AllSeatsNormalRoom)
		r.set.usedSeats(0)

	}
	//r.set.name(3)
	// if its first then? if end then
}

// ===============================================================
/*
type cat struct {
name string
}
myCat.set.name("Allen")
*/
