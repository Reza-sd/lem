package room

// ===============================================================

func (r *Room) HasOneFreeSeat() bool {

	return r.get.usedSeats() < r.get.allSeats()
}

// ===============================================================
