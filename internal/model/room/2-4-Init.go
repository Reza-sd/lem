package room

func (myRoom *Room)Init(name,allSeats Mtr , connectionSlice []Mtr){
	myRoom.Name=name
	myRoom.AllSeats=allSeats
	myRoom.UsedSeats=0
}