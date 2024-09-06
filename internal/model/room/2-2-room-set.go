package room

// ================================
func (r *room) SetName(name m) *room {
	r.name = name
	return r
}
func (r *room) SetAllSeats(allSeats m) *room {
	r.allSeats = allSeats
	return r
}
func (r *room) SetUsedSeats(usedSeats m) *room {
	r.usedSeats = usedSeats
	return r
}
func (r *room) SetConnectionSlice(connectionSlice []m) *room {
	r.connectionSlice = make([]m, len(connectionSlice), maxLenConnectionSlice)
	copy(r.connectionSlice, connectionSlice)

	return r
}

// ===============================================================
func (r *room) UpdateOneCome() (*room, []e) {

	if r.GetUsedSeats() == r.GetAllSeats() || r.GetUsedSeats()+1 > r.GetAllSeats() {
		return r, Wrapper(UpdateOneCome10, nil)
	}

	r.SetUsedSeats(r.GetUsedSeats() + 1)
	return r, nil
}

// =====================================================
func (r *room) UpdateOneLeave() (*room, []e) {
	if r.GetUsedSeats() == 0 {
		return r, Wrapper(UpdateOneLeave10, nil)
	}
	r.SetUsedSeats(r.GetUsedSeats() - 1)
	return r, nil
}

// =====================================================
