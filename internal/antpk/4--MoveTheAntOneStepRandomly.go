package antpk

import (
	//"fmt"
	//"fmt"
	"fmt"
	graphpk "main/internal/graphpk"
	//"math/rand"
	//"slices"
)

// ==========================================================
func (theAnt *Ant) MoveTheAntOneStepRandomly(theGraph *graphpk.Graph, travelHistory *TravelHistory) error {
	funcName := "MoveTheAntOneStepRandomly"
	//fmt.Println("move theAnt")
	//----------------input validation--healthy Ant----------------
	if theAnt.StepNumber != len(theAnt.VisitedRoomsArr)-1 {
		return logger.RWarn(funcName)
	}

	//---------------------------------------------------

	// -----check if its already in End room----
	if theAnt.CurrentRoomName == theGraph.EndRoomName {
		return
	}

	nextRandomAvailableRoomName, _ := theAnt.nextRandomAvailableRoomNamePassableTunnel(theGraph, travelHistory)
	// in the case of no available free next room
	if nextRandomAvailableRoomName == "" {
		theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, "*")
		theAnt.StepNumber++
		//theAnt.StepNumber=theGraph.
		//nextRoom.LastAntCameFromRoomName =""
		return
	}
	//---------------------------------------------

	//---------------------------------------------
	currentRoom := theGraph.Rooms[theAnt.CurrentRoomName]
	currentRoom.EmptySeats = 1
	theGraph.Rooms[theAnt.CurrentRoomName] = currentRoom
	fmt.Println()
	nextRoom := theGraph.Rooms[nextRandomAvailableRoomName]
	nextRoom.EmptySeats = 0

	//---------------

	if nextRandomAvailableRoomName == theGraph.EndRoomName {
		nextRoom.EmptySeats = 1000
		theGraph.CurrentAntsInEndRoom++
	}

	theGraph.Rooms[nextRandomAvailableRoomName] = nextRoom

	theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, nextRandomAvailableRoomName)
	theAnt.CurrentRoomName = nextRandomAvailableRoomName
	theAnt.StepNumber++
	//fmt.Println(theAnt)
}

// ==========================================================
