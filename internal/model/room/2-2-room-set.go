package room

// ================================
func (set *setter) name(name Mtr) *setter {
	set.room.name = name
	return set
}
func (set *setter) allSeats(allSeats Mtr) *setter {
	set.room.allSeats = allSeats
	return set
}
func (set *setter) usedSeats(usedSeats Mtr) *setter {
	set.room.usedSeats = usedSeats
	return set
}
func (set *setter) connectionSlice(connectionSlice []Mtr) *setter {
	set.room.connectionSlice = connectionSlice
	return set
}

// ===============================================================
func (set *setter) builder(name, lastRoomName Mtr, connectionSlice []Mtr) *room {
	set.name(name).connectionSlice(connectionSlice)
	if name == 0 || name == lastRoomName {
		set.allSeats(MaxSeatsStartEnd).usedSeats(UsedSeatsStartEnd)
	} else {
		set.allSeats(AllSeatsNormalRoom).usedSeats(0)

	}
	return set.room
	// if its first then? if end then
}

//===============================================================
