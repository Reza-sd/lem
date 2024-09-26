package room

// --------------------------------
func (set *roomTsetT) connectionSlice(connectionSlice []rT) []errT {
	set.room.data.connectionSlice = make([]rT, len(connectionSlice), maxLenConnectionSlice)
	copy(set.room.data.connectionSlice, connectionSlice)

	return nil
}

//-------------------------------
