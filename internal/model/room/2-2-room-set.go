package room

// ================================
func (r *room) SetName(name RT) *room {
	r.name = name
	return r
}
func (r *room) SetAllSeats(allSeats RT) *room {
	r.allSeats = allSeats
	return r
}
func (r *room) SetUsedSeats(usedSeats RT) *room {
	r.usedSeats = usedSeats
	return r
}
func (r *room) SetConnectionSlice(connectionSlice []RT) *room {
	r.connectionSlice = make([]RT, len(connectionSlice), maxLenConnectionSlice)
	copy(r.connectionSlice, connectionSlice)

	return r
}

// ===============================================================
func (r *room) UpdateOneCome() (*room, []statType) {

	if r.GetUsedSeats() == r.GetAllSeats() || r.GetUsedSeats()+1 > r.GetAllSeats() {
		return r, wrapper(UpdateOneCome10, nil)
	}

	r.SetUsedSeats(r.GetUsedSeats() + 1)
	return r, nil
}

// =====================================================
func (r *room) UpdateOneLeave() (*room, []statType) {
	if r.GetUsedSeats() == 0 {
		return r, wrapper(UpdateOneLeave10, nil)
	}
	r.SetUsedSeats(r.GetUsedSeats() - 1)
	return r, nil
}

// =====================================================
