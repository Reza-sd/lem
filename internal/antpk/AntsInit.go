package antpk

import (
	"fmt"
	//graphpk "main/internal/graphpk"
)

func (myAntGroup *AntGroup) AntsInit(numberOfAnts int, startRoomName string) error {
	//-----input validation-----
	if numberOfAnts < 1 {
		return fmt.Errorf("number of Ants < 1")
	}
	if startRoomName == "" {
		return fmt.Errorf("startRoomName is EMPTY ")
	}
	//--------------------------
	myAntGroup.AntsMap = make(map[int]Ant) //initiate map
	myAntGroup.NumberOfAnts = numberOfAnts

	for i := 1; i <= numberOfAnts; i++ {
		name := fmt.Sprintf("L%d", i) //Generate name for each Ant
		myAntGroup.AntsMap[i] = Ant{
			Name:            name,
			VisitedRoomsArr: []string{startRoomName}, CurrentRoomName: startRoomName,
			StepNumber: 0,
		}

	}
	return nil
}
