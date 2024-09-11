package room

// ================================
type rSetter struct {
	room *room
}

// ================================
func (set *rSetter) name(name RT) errArrT {
	//check if name valid to set
	set.room.data.name = name
	return nil
}
func (set *rSetter) allSeats(allSeats RT) errArrT {
	//check validation
	set.room.data.allSeats = allSeats
	return nil
}
func (set *rSetter) usedSeats(usedSeats RT) errArrT {
	set.room.data.usedSeats = usedSeats
	return nil
}
func (set *rSetter) connectionSlice(connectionSlice []RT) errArrT {
	set.room.data.connectionSlice = make([]RT, len(connectionSlice), maxLenConnectionSlice)
	copy(set.room.data.connectionSlice, connectionSlice)

	return nil
}

//================================
