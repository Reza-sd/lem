package room

// ===============================================================

func (r *room) HasOneFreeSeat() bool {

	return r.get.usedSeats() < r.get.allSeats()
}

// ===============================================================
