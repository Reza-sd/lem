package room

func (r *Room) Init(name, lastRoomName Mtr, connectionSlice []Mtr) {
	//myRoom.name = name
	r.setName(name)
	r.setConnectionSlice(connectionSlice)
	//r.connectionSlice = connectionSlice

	if name == 0 || name == lastRoomName {
		r.setAllSeats(MaxSeatsStartEnd)
		r.setUsedSeats(UsedSeatsStartEnd)
		//r.allSeats = MaxSeatsStartEnd
		//r.usedSeats = UsedSeatsStartEnd
	} else {
		r.setAllSeats(AllSeatsNormalRoom)
		//r.allSeats = AllSeatsNormalRoom
		r.setUsedSeats(0)
		//r.usedSeats = 0
	}

	// if its first then? if end then
}
