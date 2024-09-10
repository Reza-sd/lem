package room

//================hasOneFreeSeat===============================

func (get *rGetter) hasOneFreeSeat() bool {

	return get.UsedSeats() < get.AllSeats()
}

//================================================
