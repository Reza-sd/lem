package room

// ===================================
type roomTgetT struct {
	room *roomT
}

// ================================================
func (get *roomTgetT) Name() rT {
	return get.room.data.name //r.data.name
}

func (get *roomTgetT) AllSeats() rT {
	return get.room.data.allSeats
}
func (get *roomTgetT) UsedSeats() rT {
	return get.room.data.usedSeats
}
func (get *roomTgetT) ConnectionSlice() []rT {
	return get.room.data.connectionSlice

}

//===================================================
