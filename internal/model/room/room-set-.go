package room

// ================================
type rSetter struct {
	room *room
}

//-------------------------------------------------
func (set *rSetter) allSeats(allSeats RT) []errT {
	//check validation
	set.room.data.allSeats = allSeats
	return nil
}
func (set *rSetter) usedSeats(usedSeats RT) []errT {
	set.room.data.usedSeats = usedSeats
	return nil
}
func (set *rSetter) connectionSlice(connectionSlice []RT) []errT {
	set.room.data.connectionSlice = make([]RT, len(connectionSlice), maxLenConnectionSlice)
	copy(set.room.data.connectionSlice, connectionSlice)

	return nil
}

//================================
