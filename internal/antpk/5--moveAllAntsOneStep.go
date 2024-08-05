package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *AntGroup) MoveAllAntsOneStepRandomly(theGraph *graphpk.Graph) error {
	funcName := "MoveAllAntsOneStepRandomly"

	// for i := 1; i <= allAnts.NumberOfAnts; i++ {
	for _, theAntName := range allAnts.notArrivedAntsName {

		theAnt := allAnts.AntsDb[theAntName]

		if theAnt.CurrentRoomName == theGraph.EndRoomName {
			// already arrived
			continue
		}

		//canIMove, _, err := CanImoveWhere("L1", myAntGroup, myGraph)

		canIMove, moveTo, err := CanImoveWhere(theAnt.Name, *allAnts, *theGraph)

		if err != nil {
			return logger.RErr(funcName, "CanImoveWhere", err, "give next avalable room")
		}

		if canIMove == false {

		}
		//currentRoomName := theAnt.CurrentRoomName
		//theAnt.MoveTheAntOneStepRandomly(theGraph, travelHistory)
		//MovedRoomName := theAnt.CurrentRoomName

		//theGraph.UsedTunnelsInLastSequence[currentRoomName] = MovedRoomName

		//allAnts.AntsMap[i] = theAnt

	}
	return nil
}
