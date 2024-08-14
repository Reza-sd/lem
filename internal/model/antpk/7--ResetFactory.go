package antpk

// =============================================
func (myAntGroup *AntGroup) ResetFactory() {

	myAntGroup.CurrentSequenceNumber = 0
	myAntGroup.UsedTunnel = make(map[Mta][]Mta)
	myAntGroup.NotArrivedAntsName = make(map[Mta]struct{})

	for antName := range myAntGroup.AntsDb {
		myAntGroup.NotArrivedAntsName[antName] = struct{}{}
		ant := myAntGroup.AntsDb[antName]
		ant.CurrentRoomName = 0
		ant.StepNumber = 0
		ant.VisitedRoomsArr = []Mta{0}
		myAntGroup.AntsDb[antName] = ant
	}

}

//=============================================
