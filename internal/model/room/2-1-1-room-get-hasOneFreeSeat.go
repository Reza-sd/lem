package room

//================hasOneFreeSeat===============================

func (get *roomGetter) hasOneFreeSeat() bool {

	return get.UsedSeats() < get.AllSeats()
}
//================================================