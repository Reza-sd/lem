package room

// ================================
func (r *room) SetName(name mtr) *room {
	r.name = name
	return r
}
func (r *room) SetAllSeats(allSeats mtr) *room {
	r.allSeats = allSeats
	return r
}
func (r *room) SetUsedSeats(usedSeats mtr) *room {
	r.usedSeats = usedSeats
	return r
}
func (r *room) SetConnectionSlice(connectionSlice []mtr) *room {
	r.connectionSlice = make([]mtr, len(connectionSlice), maxLenConnectionSlice)
	copy(r.connectionSlice, connectionSlice)

	return r
}

// ===============================================================
func (r *room) UpdateOneCome() *room {

	if r.GetUsedSeats() == r.GetAllSeats() || r.GetUsedSeats()+1 > r.GetAllSeats() {
		return Answer[*room](r, Room_set_oneCome, Room_set_oneCome_code_10, r)
	}

	r.SetUsedSeats(r.GetUsedSeats() + 1)
	return Answer[*room](r, Room_set_oneCome, null, r)
}

// =====================================================
func (r *room) UpdateOneLeave() *room {
	if r.GetUsedSeats() == 0 {
		return Answer[*room](r, Room_set_oneLeave, Room_set_oneLeave_code_10, r)
	}
	r.SetUsedSeats(r.GetUsedSeats() - 1)
	return Answer[*room](r, Room_set_oneLeave, null, r)
}

// =====================================================
