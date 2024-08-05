package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *AntGroup) TryMoveAllAntsOneStepRandomly(theGraph *graphpk.Graph) error {

	funcName := "MoveAllAntsOneStepRandomly"

	//-------------------------------
	if len(allAnts.NotArrivedAntsName) == 0 {
		return nil
	}
	//----------------------------
	allAnts.CurrentSequenceNumber++
	//----------------------------------
	for i := 1; i < 50; i++ { // try

		//for index, theAntName := range allAnts.NotArrivedAntsName {
		for theAntName := range allAnts.NotArrivedAntsName {

			theAnt, okTheAnt := allAnts.AntsDb[theAntName]

			if !okTheAnt {
				return logger.RWarnStr(funcName, "AntsDb", "the ant does not exist", "")
			}

			if theAnt.CurrentRoomName == theGraph.EndRoomName {
				// already arrived
				continue
			}

			if theAnt.StepNumber == allAnts.CurrentSequenceNumber {
				continue
			}

			canIMove, moveTo, err := CanImoveWhere(theAnt.Name, *allAnts, *theGraph)

			if err != nil {
				return logger.RErr(funcName, "CanImoveWhere", err, "give next avalable room")
			}

			//if !canIMove {
			//"*"
			//theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, "*")
			//theAnt.StepNumber++
			//allAnts.currentSequenceNumber=theAnt.StepNumber
			//}

			if canIMove {

				seq := allAnts.CurrentSequenceNumber
				from := theAnt.CurrentRoomName
				// _ = seq
				// _ = from
				allAnts.UsedTunnel[seq] = make(map[string]string) //Initialize
				allAnts.UsedTunnel[seq][from] = moveTo

				theAnt.CurrentRoomName = moveTo
				theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, moveTo)
				theAnt.StepNumber++

				moveToRoomObject := theGraph.Rooms[moveTo]
				moveToRoomObject.EmptySeats--
				theGraph.Rooms[moveTo] = moveToRoomObject

				if theAnt.CurrentRoomName == theGraph.EndRoomName {
					//allAnts.NotArrivedAntsName = append(allAnts.NotArrivedAntsName[:index], allAnts.NotArrivedAntsName[index+1:]...)
					delete(allAnts.NotArrivedAntsName, theAnt.CurrentRoomName)
				}

			}

			allAnts.AntsDb[theAntName] = theAnt

		}
	}
	//---------------
	//for _,theAntName :=range allAnts.NotArrivedAntsName {
	for theAntName := range allAnts.NotArrivedAntsName {

		theAnt, okTheAnt := allAnts.AntsDb[theAntName]

		if !okTheAnt {
			return logger.RWarnStr(funcName, "AntsDb", "the ant does not exist", "")
		}

		if theAnt.StepNumber != allAnts.CurrentSequenceNumber {
			theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, "*")
			theAnt.StepNumber++
			allAnts.AntsDb[theAntName] = theAnt

		}

		if theAnt.CurrentRoomName == theGraph.EndRoomName {
			delete(allAnts.NotArrivedAntsName, theAntName)
		}

	}

	//----------------------
	return nil
}
