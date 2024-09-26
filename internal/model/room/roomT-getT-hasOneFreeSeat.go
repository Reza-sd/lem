package room

//========hasOneFreeSeat=====================

func (get *getT) hasOneFreeSeat() bool {

	return get.UsedSeats() < get.AllSeats()
}

//================================================
