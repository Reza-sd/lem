package room

func (r *Room) HasOneFreeSeat() bool {

	// if myRoom.usedSeats<myRoom.AllSeats {
	// 	return true
	// }
	//return r.usedSeats < r.allSeats
	return r.getUsedSeats() < r.getAllSeats()
}
