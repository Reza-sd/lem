package antpk

//"fmt"

func (myAntGroup *AntGroup) Init(numberOfAnts Mtag) error {
	funcName := "AntsInit"
	//var err error
	//-----input validation-----
	if numberOfAnts < 1 || numberOfAnts > MaxHandleableAntsNumber {
		return logger.RWarnStr(funcName, "numberOfAnts[1-max]", "is not valid", "input validation numberOfAnts")
	}
	//--------------------------
	myAntGroup.AntsDb = make(map[Mtag]Ant) //initiate map
	myAntGroup.NumberOfAnts = numberOfAnts
	myAntGroup.CurrentSequenceNumber = 0
	myAntGroup.UsedTunnel = make(map[Mtag][]Mtag)
	myAntGroup.NotArrivedAntsName = make(map[Mtag]struct{})
	//--------------------------
	//myAntGroup
	for i := Mtag(1); i <= numberOfAnts; i++ {
		myAntGroup.AntsDb[i] = Ant{
			//Name: i,

			CurrentRoomName: 0,
			VisitedRoomsArr: []Mtag{0},
			StepNumber:      0,
		}

		myAntGroup.NotArrivedAntsName[i] = struct{}{}

	}
	return nil
}
