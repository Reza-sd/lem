package antpk

import (
	"fmt"
	graphpk "main/internal/graphpk"
)

func (myAntGroup *AntGroup) AntsInit(numberOfAnts int, myGraph *graphpk.Graph) {

	myAntGroup.AntsMap = make(map[int]Ant)
	myAntGroup.NumberOfAnts = numberOfAnts
	//myAnts.UsedTunnelsInLastSequence=[]string{}

	for i := 1; i <= numberOfAnts; i++ {
		name := fmt.Sprintf("L%d", i) //Generate name for each Ant
		myAntGroup.AntsMap[i] = Ant{
			Name:            name,
			VisitedRoomsArr: []string{myGraph.StartRoomName}, CurrentRoomName: myGraph.StartRoomName,
			StepNumber: 0,
		}

	}

}
