package move

import (
	"fmt"
	data "main/internal/data"
)

func (myAnts *Ants) AntsInit(num int, myGraph *data.Graph) {

	myAnts.Ants = make(map[int]Ant)

	for i := 1; i <= num; i++ {
		name := fmt.Sprintf("L%d", i)
		myAnts.Ants[i] = Ant{Name: name, currentRoomName: myGraph.StartRoomName}

	}

}
