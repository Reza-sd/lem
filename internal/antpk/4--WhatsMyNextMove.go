package antpk

import (
	//"fmt"
	graphpk "main/internal/graphpk"
	"math/rand"
	"slices"
)

// =====================================================
func WhatsMyNextMove(antName string, theAntGroup AntGroup, theGraph graphpk.Graph) (string, error) {
	funcName := "WhatsMyNextMove"
	theAnt, okAnt := theAntGroup.AntsDb[antName]
	if !okAnt {
		return "", logger.RWarnStr(funcName, "okAnt", "the AntGroup does not have this Ant", "check if ant avaliable in antsgroup")
	}
	//Am I able to move?
	if theAnt.CurrentRoomName == theGraph.EndRoomName {
		//already arrived
		return "*", nil //means stop
	}
	//-------------------
	currentRoomObjectFromGraph, okRoom := theGraph.Rooms[theAnt.CurrentRoomName]
	if !okRoom {
		return "", logger.RWarnStr(funcName, "okRoom", "the graph does not have this room", "check if graph has this room name")
	}
	//----------------
	currentConnectionsArr := currentRoomObjectFromGraph.Connections
	lengthcurrentConnectionsArr := len(currentConnectionsArr)
	//------------
	for i := 0; i < 50; i++ {
		randomNextRoomIndex := rand.Intn(lengthcurrentConnectionsArr)
		nextRandomAvailableRoomName := currentConnectionsArr[randomNextRoomIndex]
		//---check if I visited before---------
		if slices.Contains(theAnt.VisitedRoomsArr, nextRandomAvailableRoomName) {
			continue
		}
		//---to check if the offered tunnel(from,to) is used in this sequence----
		UsedTunnelMap, ok := theAntGroup.UsedTunnel[theAntGroup.currentSequenceNumber]
		if ok {
			if UsedTunnelMap[theAnt.CurrentRoomName] == nextRandomAvailableRoomName {
				continue
			}
		}
		//---------------------
		nextRoom := theGraph.Rooms[nextRandomAvailableRoomName]
		// check if the next selected room empty seat
		if nextRoom.EmptySeats == 0 {
			continue
		}
		//-----------------------------
		return nextRandomAvailableRoomName, nil
	}
	//----------------------------------
	return "*", nil //means stop

}

// ==========================================================
// func (theAnt *Ant) MoveTheAntOneStepRandomly(theGraph *graphpk.Graph, travelHistory *TravelHistory) error {
// 	funcName := "MoveTheAntOneStepRandomly"
// 	//fmt.Println("move theAnt")
// 	// the first question : Am I able to move?
// 	//----------------input validation--healthy Ant----------------
// 	if theAnt.StepNumber != len(theAnt.VisitedRoomsArr)-1 {
// 		return logger.RWarn(funcName)
// 	}

// 	//---------------------------------------------------

// 	// -----check if its already in End room----
// 	if theAnt.CurrentRoomName == theGraph.EndRoomName {
// 		return
// 	}

// 	nextRandomAvailableRoomName, _ := theAnt.nextRandomAvailableRoomNamePassableTunnel(theGraph, travelHistory)
// 	// in the case of no available free next room
// 	if nextRandomAvailableRoomName == "" {
// 		theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, "*")
// 		theAnt.StepNumber++
// 		//theAnt.StepNumber=theGraph.
// 		//nextRoom.LastAntCameFromRoomName =""
// 		return
// 	}
// 	//---------------------------------------------

// 	//---------------------------------------------
// 	currentRoom := theGraph.Rooms[theAnt.CurrentRoomName]
// 	currentRoom.EmptySeats = 1
// 	theGraph.Rooms[theAnt.CurrentRoomName] = currentRoom
// 	fmt.Println()
// 	nextRoom := theGraph.Rooms[nextRandomAvailableRoomName]
// 	nextRoom.EmptySeats = 0

// 	//---------------

// 	if nextRandomAvailableRoomName == theGraph.EndRoomName {
// 		nextRoom.EmptySeats = 1000
// 		theGraph.CurrentAntsInEndRoom++
// 	}

// 	theGraph.Rooms[nextRandomAvailableRoomName] = nextRoom

// 	theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, nextRandomAvailableRoomName)
// 	theAnt.CurrentRoomName = nextRandomAvailableRoomName
// 	theAnt.StepNumber++
// 	//fmt.Println(theAnt)
// }

// ==========================================================
