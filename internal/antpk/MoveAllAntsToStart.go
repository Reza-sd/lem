package antpk

import (
	//"fmt"
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *Ants) MoveAllAntsToStart(theGraph *graphpk.Graph) {

	// for i := 1; i <= allAnts.AntsNumber; i++ {
	// 	theAnt := allAnts.AntsMap[i]
	//     theAnt.CurrentRoomName=theGraph.StartRoomName
	// 	theAnt.VisitedRoomsArr = []string{theGraph.StartRoomName}
	//     //fmt.Println(theAnt)
	// }

	for _, ant := range allAnts.AntsMap {
		ant.CurrentRoomName = theGraph.StartRoomName
		ant.VisitedRoomsArr = []string{theGraph.StartRoomName}
	}
}
