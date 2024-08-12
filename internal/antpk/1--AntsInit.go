package antpk

//"fmt"

func (myAntGroup *AntGroup) AntGroupInit(numberOfAnts mt) error {
	funcName := "AntsInit"
	//var err error
	//-----input validation-----
	if numberOfAnts < 1 || numberOfAnts > MaxHandleableAntsNumber {
		return logger.RWarnStr(funcName, "numberOfAnts[1-max]", "is not valid", "input validation numberOfAnts")
	}
	//--------------------------
	myAntGroup.AntsDb = make(map[mt]Ant) //initiate map
	myAntGroup.NumberOfAnts = numberOfAnts
	myAntGroup.CurrentSequenceNumber = 0
	myAntGroup.UsedTunnel = make(map[mt][]mt)
	myAntGroup.NotArrivedAntsName = make(map[mt]struct{})
	//--------------------------
	//myAntGroup
	for i := mt(1); i <= numberOfAnts; i++ {
		myAntGroup.AntsDb[i] = Ant{
			Name: i,

			CurrentRoomName: 0,
			VisitedRoomsArr: []mt{0},
			StepNumber:      0,
		}

		myAntGroup.NotArrivedAntsName[i] = struct{}{}

	}
	return nil
}
