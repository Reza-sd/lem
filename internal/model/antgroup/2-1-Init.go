package antgroup

//"fmt"

func (myAntGroup *AntGroup) Init(numberOfAnts Mtag) error {
	funcName := "AntsInit"
	//var err error
	//-----input validation-----
	if numberOfAnts < 1 || numberOfAnts > MaxHandleableAntsNumber {
		return logger.WarnLogRErrMsg(funcName, "numberOfAnts[1-max]", "is not valid", "input validation numberOfAnts")
	}
	//--------------------------
	myAntGroup.AntsDb = make(map[Mtag]*Ant) //initiate map
	myAntGroup.NumberOfAnts = numberOfAnts
	myAntGroup.CurrentSequenceNumber = 0
	myAntGroup.UsedTunnel = make(map[Mtag][]Mtag)
	myAntGroup.NotArrivedAntsName = make(map[Mtag]struct{})
	//--------------------------
	
	for i := Mtag(1); i <= numberOfAnts; i++ {
	
		//var myAnt Ant
		//myAnt := new(Ant) //return a pointer to new instance
		myAnt := &Ant{}
		myAnt.Init(i)
		myAntGroup.AntsDb[i] = myAnt

		myAntGroup.NotArrivedAntsName[i] = struct{}{}

	}
	return nil
}
