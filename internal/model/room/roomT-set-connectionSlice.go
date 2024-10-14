package room

// --------------------------------
func (set *roomTsetT) connectionSlice(connectionSlice []rT) []errT {

	set.room.data.connectionSlice = make([]rT, len(connectionSlice), _MAX_LEN_CONNECTION_SLICE)

	copy(set.room.data.connectionSlice, connectionSlice)

	return nil
}

//-------------------------------
