package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *AntGroup) MoveAllAntsOneStepRandomly(theGraph *graphpk.Graph, travelHistory *TravelHistory) {

	for i := 1; i <= allAnts.NumberOfAnts; i++ {
		theAnt := allAnts.AntsMap[i]
		if theAnt.CurrentRoomName == theGraph.EndRoomName {
			// already arrived
			continue
		}

		//currentRoomName := theAnt.CurrentRoomName
		theAnt.MoveTheAntOneStepRandomly(theGraph, travelHistory)
		//MovedRoomName := theAnt.CurrentRoomName

		//theGraph.UsedTunnelsInLastSequence[currentRoomName] = MovedRoomName

		allAnts.AntsMap[i] = theAnt

	}

}
