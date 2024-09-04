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
	set.room.connectionSlice = make([]Mtr, len(connectionSlice))
	copy(set.room.connectionSlice, connectionSlice)

	return set
}

// ===============================================================
