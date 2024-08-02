package antpk

import (
	"fmt"
)

func (myAntGroup *AntGroup) AntsInit(numberOfAnts int, startRoomName string) error {
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
	myAntGroup.AntsMap = make(map[int]Ant) //initiate map
	myAntGroup.NumberOfAnts = numberOfAnts
	myAntGroup.SequenceNumber = 0

	for i := 1; i <= numberOfAnts; i++ {
		name := fmt.Sprintf("L%d", i) //Generate name for each Ant
		myAntGroup.AntsMap[i] = Ant{
			Name: name,

			CurrentRoomName: startRoomName,
			VisitedRoomsArr: []string{startRoomName},
			StepNumber:      0,
		}

	}
	return nil
}
