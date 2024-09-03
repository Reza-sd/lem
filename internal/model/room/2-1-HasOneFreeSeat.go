package room

// ===============================================================

func (r *Room) HasOneFreeSeat() bool {

	return r.getUsedSeats() < r.getAllSeats()
}

// ===============================================================
