package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *AntGroup) MoveAllAntsOneStepRandomly(theGraph *graphpk.Graph) error {

	funcName := "MoveAllAntsOneStepRandomly"

	for _, theAntName := range allAnts.NotArrivedAntsName {

		theAnt,okTheAnt := allAnts.AntsDb[theAntName]

		if !okTheAnt{
			 return logger.RWarnStr(funcName, "AntsDb", "the ant does not exist","")
		}

		if theAnt.CurrentRoomName == theGraph.EndRoomName {
			// already arrived
			continue
		}

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
			_=seq
			_=from
			allAnts.UsedTunnel[seq][from] = moveTo

			theAnt.CurrentRoomName = moveTo
			theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, moveTo)
			theAnt.StepNumber++

			moveToRoomObject := theGraph.Rooms[moveTo]
			moveToRoomObject.EmptySeats--
			theGraph.Rooms[moveTo] = moveToRoomObject
		}

		allAnts.AntsDb[theAntName] = theAnt

	}

	allAnts.CurrentSequenceNumber++
	return nil
}
