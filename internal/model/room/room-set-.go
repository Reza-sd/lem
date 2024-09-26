package room

// ================================
type setT struct {
	room *room
}

// -------------------------------------------------
func (set *setT) allSeats(allSeats rT) []errT {
	//check validation
	set.room.data.allSeats = allSeats
	return nil
}
func (set *setT) usedSeats(usedSeats rT) []errT {
	set.room.data.usedSeats = usedSeats
	return nil
}
func (set *setT) connectionSlice(connectionSlice []rT) []errT {
	set.room.data.connectionSlice = make([]rT, len(connectionSlice), maxLenConnectionSlice)
	copy(set.room.data.connectionSlice, connectionSlice)

	return nil
}

//================================
