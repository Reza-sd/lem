package room

// ==================================================
func (get *rGetter) Name() RT {
	return get.room.data.name //r.data.name
}

func (get *rGetter) AllSeats() RT {
	return get.room.data.allSeats
}
func (get *rGetter) UsedSeats() RT {
	return get.room.data.usedSeats
}
func (get *rGetter) ConnectionSlice() []RT {
	return get.room.data.connectionSlice

}

// ==========================================================
