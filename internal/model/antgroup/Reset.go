package antpk

// =============================================
func (myAntGroup *AntGroup) Reset() {

	myAntGroup.CurrentSequenceNumber = 0
	myAntGroup.UsedTunnel = make(map[Mtag][]Mtag)
	myAntGroup.NotArrivedAntsName = make(map[Mtag]struct{})

	for antName := range myAntGroup.AntsDb {
		myAntGroup.NotArrivedAntsName[antName] = struct{}{}
		ant := myAntGroup.AntsDb[antName]
		ant.CurrentRoomName = 0
		ant.StepNumber = 0
		ant.VisitedRoomsArr = []Mtag{0}
		myAntGroup.AntsDb[antName] = ant
	}

}

//=============================================
