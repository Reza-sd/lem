package room

// ===============================================================
func (r *room) initiator(name, lastRoomName Mtr, connectionSlice []Mtr) {
	r.set.name(name).connectionSlice(connectionSlice)
	if name == 0 || name == lastRoomName {
		r.set.allSeats(MaxSeatsStartEnd).usedSeats(UsedSeatsStartEnd)
	} else {
		r.set.allSeats(AllSeatsNormalRoom).usedSeats(0)

	}

	// if its first then? if end then
}

//===============================================================
