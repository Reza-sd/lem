package room

// ===============================================================
func (r *Room) initiator(name, lastRoomName Mtr, connectionSlice []Mtr) {
	r.setName(name)
	r.setConnectionSlice(connectionSlice)
	if name == 0 || name == lastRoomName {
		r.setAllSeats(MaxSeatsStartEnd)
		r.setUsedSeats(UsedSeatsStartEnd)
	} else {
		r.setAllSeats(AllSeatsNormalRoom)
		r.setUsedSeats(0)

	}

	// if its first then? if end then
}

// ===============================================================
