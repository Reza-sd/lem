package antpk

import (
	"fmt"
)

func (myAntGroup *AntGroup) AntGroupInit(numberOfAnts int, startRoomName string) error {
	funcName := "AntsInit"
	//var err error
	//-----input validation-----
	if numberOfAnts < 1 || numberOfAnts > MaxHandleableAntsNumber {
		return logger.RWarnStr(funcName, "numberOfAnts[1-max]", "is not valid", "input validation numberOfAnts")
	}
	//-------
	if startRoomName == "" {
		return logger.RWarnStr(funcName, "startRoomName ? EMPTY", "is not valid", "check if startRoomName not empty")
	}
	//--------------------------
	myAntGroup.AntsDb = make(map[string]Ant) //initiate map
	myAntGroup.NumberOfAnts = numberOfAnts
	myAntGroup.CurrentSequenceNumber = 0
	myAntGroup.UsedTunnel = map[int]map[string]string{}
	//--------------------------
	for i := 1; i <= numberOfAnts; i++ {
		name := fmt.Sprintf("L%d", i) //Generate name for each Ant
		myAntGroup.AntsDb[name] = Ant{
			Name: name,

			CurrentRoomName: startRoomName,
			VisitedRoomsArr: []string{startRoomName},
			StepNumber:      0,
		}
		myAntGroup.NotArrivedAntsName = append(myAntGroup.NotArrivedAntsName, name)

	}
	return nil
}
