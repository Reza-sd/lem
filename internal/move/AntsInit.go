package move

import (
	"fmt"
	data "main/internal/data"
)

func (myAnts *Ants) AntsInit(num int, myGraph *data.Graph) {

	myAnts.Ants = make(map[int]Ant)
	//myAnts.AntsNumber = num

	fisrtVisited := []string{myGraph.StartRoomName}

	for i := 1; i <= num; i++ {
		name := fmt.Sprintf("L%d", i)
		myAnts.Ants[i] = Ant{Name: name, VisitedRooms: fisrtVisited, CurrentRoomName: myGraph.StartRoomName}

	}
	//-----

	//------

}
