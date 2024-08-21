package ant

func (myAnt *Ant) Init(antName Mta) {

	myAnt.Name = antName
	myAnt.CurrentRoomName = 0
	myAnt.StepNumber = 0
	myAnt.VisitedRoomsArr = []Mta{0}
}
