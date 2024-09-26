package room

//========hasOneFreeSeat=====================

func (get *roomTgetT) hasOneFreeSeat() bool {

	return get.UsedSeats() < get.AllSeats()
}

//================================================
