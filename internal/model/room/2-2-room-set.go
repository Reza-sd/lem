package room

// ================================
func (set *setter) name(name mtr) *setter {
	set.room.name = name
	return set
}
func (set *setter) allSeats(allSeats mtr) *setter {
	set.room.allSeats = allSeats
	return set
}
func (set *setter) usedSeats(usedSeats mtr) *setter {
	set.room.usedSeats = usedSeats
	return set
}
func (set *setter) connectionSlice(connectionSlice []mtr) *setter {
	set.room.connectionSlice = make([]mtr, len(connectionSlice))
	copy(set.room.connectionSlice, connectionSlice)

	return set
}

// ===============================================================
