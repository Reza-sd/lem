package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *Ants) MoveAllAntsOneStepRandomly(theGraph *graphpk.Graph) {
	//stepCounter :=0

	for i := 1; i <= allAnts.NumberOfAnts; i++ {
		theAnt := allAnts.AntsMap[i]
		if theAnt.CurrentRoomName == theGraph.EndRoomName {
			// already arrived
			continue
		}
		//--------------------------------------
		//lastUsedTunnel

		//-----------------------------------------
		theAnt.MoveOneStepForwardRandomly(theGraph)

		allAnts.AntsMap[i] = theAnt
	}

	//allAnts.Step++
}
