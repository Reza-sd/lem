package antpk

//=============================================
func (myAntGroup *AntGroup) ResetFactory() {
	
	myAntGroup.CurrentSequenceNumber = 0
	myAntGroup.UsedTunnel = make(map[mt][]mt)
	myAntGroup.NotArrivedAntsName = make(map[mt]struct{})


	for antName := range myAntGroup.AntsDb {
		myAntGroup.NotArrivedAntsName[antName] = struct{}{}
		ant := myAntGroup.AntsDb[antName]
		ant.CurrentRoomName = 0
		ant.StepNumber = 0
		ant.VisitedRoomsArr = []mt{0}
		myAntGroup.AntsDb[antName] = ant
	}

}
//=============================================