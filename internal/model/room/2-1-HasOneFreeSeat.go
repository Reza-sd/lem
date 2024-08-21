package room

func (myRoom *Room) HasOneFreeSeat() bool {

	// if myRoom.usedSeats<myRoom.AllSeats {
	// 	return true
	// }
	return myRoom.UsedSeats < myRoom.AllSeats
}
