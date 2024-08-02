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
func (theAnt *Ant) MoveTheAntOneStepRandomly(theGraph *graphpk.Graph) {
	//fmt.Println("move theAnt")

	// -----check if its already in End room----
	if theAnt.CurrentRoomName == theGraph.EndRoomName {
		return
	}

	nextRandomAvailableRoomName := theAnt.nextRandomAvailableRoomName(theGraph)
	// in the case of no available free next room
	if nextRandomAvailableRoomName == "" {
		theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, "*")
		theAnt.StepNumber++
		//theAnt.StepNumber=theGraph.
		//nextRoom.LastAntCameFromRoomName =""
		return
	}
	//---------------------------------------------
	v, ok := theGraph.UsedTunnelsInLastSequence[theAnt.CurrentRoomName]
	if ok && v == nextRandomAvailableRoomName {

		//fmt.Println("mio")
		theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, "*")
		theAnt.StepNumber++
		delete(theGraph.UsedTunnelsInLastSequence, theAnt.CurrentRoomName)
		fmt.Println("d-UsedTunnelsInLastSequence", theGraph.UsedTunnelsInLastSequence)
		return
	}
	//Tunnel := currentRoomName + "-" + MovedRoomName
	//OfferedTunnel :=theAnt.CurrentRoomName+"-"+nextRandomAvailableRoomName
	// if slices.Contains(theGraph.UsedTunnelsInLastSequence,OfferedTunnel){

	// 	fmt.Println(OfferedTunnel)
	// 	fmt.Println(theGraph.UsedTunnelsInLastSequence)

	// 	theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, "*")

	// 	theAnt.NumberOfSteps++
	// 	//nextRoom.LastAntCameFromRoomName =""
	// 	return

	// }
	// if theAnt.CurrentRoomName ==  theGraph.startroom && nextRandomAvailableRoomName== theGraph.EndRoom ??
	//Used Tunnel in this sequence

	//---------------------------------------------
	currentRoom := theGraph.Rooms[theAnt.CurrentRoomName]
	currentRoom.EmptySeats = 1
	theGraph.Rooms[theAnt.CurrentRoomName] = currentRoom
	fmt.Println()
	nextRoom := theGraph.Rooms[nextRandomAvailableRoomName]
	nextRoom.EmptySeats = 0

	//---------------
	// nextRoom.LastAntCameFromRoomName = theAnt.CurrentRoomName

	// fmt.Println("nextRoom.LastAntCameFromRoomName==", nextRoom.LastAntCameFromRoomName, theAnt.Name)
	//------------------------
	//for the End node (has free seats always)
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
