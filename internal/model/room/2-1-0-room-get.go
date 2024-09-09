package room



// ==================================================
func (get *roomGetter) Name() RT {
	return get.room.data.name //r.data.name
}

func (get *roomGetter) AllSeats() RT {
	return get.room.data.allSeats
}
func (get *roomGetter) UsedSeats() RT {
	return get.room.data.usedSeats
}
func (get *roomGetter) ConnectionSlice() []RT {
	return get.room.data.connectionSlice
	
}

// ==========================================================


