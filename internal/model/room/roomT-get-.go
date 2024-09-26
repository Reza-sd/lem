package room

// ===================================
type getT struct {
	room *roomT
}

// ================================================
func (get *getT) Name() rT {
	return get.room.data.name //r.data.name
}

func (get *getT) AllSeats() rT {
	return get.room.data.allSeats
}
func (get *getT) UsedSeats() rT {
	return get.room.data.usedSeats
}
func (get *getT) ConnectionSlice() []rT {
	return get.room.data.connectionSlice

}

//===================================================
