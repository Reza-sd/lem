package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *AntGroup) MoveAllAntsOneStepRandomly(theGraph *graphpk.Graph) error {
	funcName := "MoveAllAntsOneStepRandomly"

	// for i := 1; i <= allAnts.NumberOfAnts; i++ {
	for _, theAntName := range allAnts.NotArrivedAntsName {

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

		if !canIMove {
			//"*"
			theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, "*")
			theAnt.StepNumber++
			//allAnts.currentSequenceNumber=theAnt.StepNumber
		} else {
			seq := allAnts.CurrentSequenceNumber
			from := theAnt.CurrentRoomName

			allAnts.UsedTunnel[seq][from] = moveTo

			theAnt.CurrentRoomName = moveTo
			theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, moveTo)
			theAnt.StepNumber++

			moveToRoomObject := theGraph.Rooms[moveTo]
			moveToRoomObject.EmptySeats = 0
			theGraph.Rooms[moveTo] = moveToRoomObject
		}
		//currentRoomName := theAnt.CurrentRoomName
		//theAnt.MoveTheAntOneStepRandomly(theGraph, travelHistory)
		//MovedRoomName := theAnt.CurrentRoomName

		//theGraph.UsedTunnelsInLastSequence[currentRoomName] = MovedRoomName

		//allAnts.AntsMap[i] = theAnt
		allAnts.AntsDb[theAntName] = theAnt

	}
	//allAnts.currentSequenceNumber++
	return nil
}
