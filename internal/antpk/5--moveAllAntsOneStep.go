package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *AntGroup) MoveAllAntsOneStepRandomly(theGraph *graphpk.Graph) {
	//stepCounter :=0
	//listOfUsedTunnelInThisSequence
	for i := 1; i <= allAnts.NumberOfAnts; i++ {
		theAnt := allAnts.AntsMap[i]
		if theAnt.CurrentRoomName == theGraph.EndRoomName {
			// already arrived
			continue
		}
		//--------------------------------------
		//lastUsedTunnel

		//-----------------------------------------
		currentRoomName := theAnt.CurrentRoomName
		theAnt.MoveTheAntOneStepRandomly(theGraph)
		MovedRoomName := theAnt.CurrentRoomName

		//Tunnel := currentRoomName + "-" + MovedRoomName

		theGraph.UsedTunnelsInLastSequence[currentRoomName] = MovedRoomName
		//append(theGraph.UsedTunnelsInLastSequence, Tunnel)
		allAnts.AntsMap[i] = theAnt

	}

	//allAnts.Step++
}
