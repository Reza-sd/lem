package ant

func (myAnt *Ant) Reset() {
	myAnt.CurrentRoomName = 0
	myAnt.StepNumber = 0
	myAnt.VisitedRoomsArr = []Mta{0}
}
