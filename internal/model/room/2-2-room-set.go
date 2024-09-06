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
func (r *room) UpdateOneCome() *room {

	if r.GetUsedSeats() == r.GetAllSeats() || r.GetUsedSeats()+1 > r.GetAllSeats() {
		return Answer[*room](r, UpdateOneCome, Room_set_oneCome_code_10, r)
	}

	r.SetUsedSeats(r.GetUsedSeats() + 1)
	return Answer[*room](r, Null, Null, r)
}

// =====================================================
func (r *room) UpdateOneLeave() *room {
	if r.GetUsedSeats() == 0 {
		return Answer[*room](r, UpdateOneLeave, UpdateOneLeave10, r)
	}
	r.SetUsedSeats(r.GetUsedSeats() - 1)
	return Answer[*room](r, Null, Null, r)
}

// =====================================================
