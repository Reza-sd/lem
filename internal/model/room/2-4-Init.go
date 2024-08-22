package room

func (myRoom *Room) Init(name, lastRoomName Mtr, connectionSlice []Mtr) {
	myRoom.Name = name
	myRoom.ConnectionSlice = connectionSlice
	if name == 0 || name == lastRoomName {
		myRoom.AllSeats = 5000
		myRoom.UsedSeats = 1000
	} else {
		myRoom.AllSeats = 1
		myRoom.UsedSeats = 0
	}

	// if its first then? if end then
}
