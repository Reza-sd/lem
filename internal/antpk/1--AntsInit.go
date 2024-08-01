package antpk

import (
	"fmt"
	//graphpk "main/internal/graphpk"
)

func (myAntGroup *AntGroup) AntsInit(numberOfAnts int, startRoomName string) error {
	//-----input validation-----
	if numberOfAnts < 1 || numberOfAnts > MaxHandleableAntsNumber {
		return fmt.Errorf("number of Ants is not in [1-200]")
	}
	if startRoomName == "" {
		return fmt.Errorf("startRoomName is EMPTY ")
	}
	//--------------------------
	myAntGroup.AntsMap = make(map[int]Ant) //initiate map
	myAntGroup.NumberOfAnts = numberOfAnts
	myAntGroup.NumberOfSequence = 0

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
