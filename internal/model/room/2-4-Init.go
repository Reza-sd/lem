package room

func (myRoom *Room) Init(name, lastRoomName Mtr, connectionSlice []Mtr) {
	myRoom.name = name
	myRoom.connectionSlice = connectionSlice
	if name == 0 || name == lastRoomName {
		myRoom.allSeats = MaxSeatsStartEnd
		myRoom.usedSeats = UsedSeatsStartEnd
	} else {
		myRoom.allSeats = AllSeatsNormalRoom
		myRoom.usedSeats = 0
	}

	// if its first then? if end then
}
