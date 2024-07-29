package move

import (
	"fmt"
	data "main/internal/data"
)

func (myAnts *Ants) AntsInit(numberOfAnts int, myGraph *data.Graph) {

	myAnts.AntsMap = make(map[int]Ant)
	myAnts.AntsNumber = numberOfAnts

	firstVisited := []string{myGraph.StartRoomName}

	for i := 1; i <= numberOfAnts; i++ {
		name := fmt.Sprintf("L%d", i) //Generate name for each Ant
		myAnts.AntsMap[i] = Ant{Name: name, VisitedRoomsArr: firstVisited, CurrentRoomName: myGraph.StartRoomName}

	}
	//-----

	//------

}
