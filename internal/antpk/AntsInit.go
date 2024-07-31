package antpk

import (
	"fmt"
	graphpk "main/internal/graphpk"
)

func (myAnts *Ants) AntsInit(numberOfAnts int, myGraph *graphpk.Graph) {

	myAnts.AntsMap = make(map[int]Ant)
	myAnts.NumberOfAnts = numberOfAnts
	//myAnts.UsedTunnelsInLastSequence=[]string{}

	for i := 1; i <= numberOfAnts; i++ {
		name := fmt.Sprintf("L%d", i) //Generate name for each Ant
		myAnts.AntsMap[i] = Ant{
			Name:            name,
			VisitedRoomsArr: []string{myGraph.StartRoomName}, CurrentRoomName: myGraph.StartRoomName,
		}

	}

}
