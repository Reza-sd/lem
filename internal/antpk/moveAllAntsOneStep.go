package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *Ants) MoveAllAntsOneStepRandomly(theGraph *graphpk.Graph) {

	for i := 1; i <= allAnts.AntsNumber; i++ {
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

}
