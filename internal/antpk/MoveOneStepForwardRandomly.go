package antpk

import (
	//"fmt"
	"fmt"
	graphpk "main/internal/graphpk"
	"math/rand"
	"slices"
)

// ------------------------------------------------
func (theAnt *Ant) MoveOneStepForwardRandomly(theGraph *graphpk.Graph) {
	//fmt.Println("move theAnt")

	// -----check if its already in End room----
	if theAnt.CurrentRoomName == theGraph.EndRoomName {
		return
	}

	nextRandomAvailableRoomName := nextRandomAvailableRoomName(theAnt, theGraph)
	// in the case of no available free next room
	if nextRandomAvailableRoomName == "" {
		theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, "*")
		return
	}

	// if theAnt.CurrentRoomName ==  theGraph.startroom && nextRandomAvailableRoomName== theGraph.EndRoom ??

	currentRoom := theGraph.Rooms[theAnt.CurrentRoomName]
	currentRoom.EmptySeats = 1
	theGraph.Rooms[theAnt.CurrentRoomName] = currentRoom
	//fmt.Println()
	nextRoom := theGraph.Rooms[nextRandomAvailableRoomName]
	nextRoom.EmptySeats = 0

	//---------------
	nextRoom.LastAntCameFromRoomName = theAnt.CurrentRoomName

	fmt.Println("nextRoom.LastAntCameFromRoomName==", nextRoom.LastAntCameFromRoomName, theAnt.Name)
	//------------------------
	//for the End node (has free seats always)
	if nextRandomAvailableRoomName == theGraph.EndRoomName {
		nextRoom.EmptySeats = 1000
		theGraph.CurrentAntsInEndRoom++
	}

	theGraph.Rooms[nextRandomAvailableRoomName] = nextRoom

	theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, nextRandomAvailableRoomName)
	theAnt.CurrentRoomName = nextRandomAvailableRoomName
	//fmt.Println(theAnt)
}

// ------------------------------------------------
func nextRandomAvailableRoomName(theAnt *Ant, theGraph *graphpk.Graph) string {
	currentRoom := theGraph.Rooms[theAnt.CurrentRoomName]
	currentTunnerArr := currentRoom.Tunnels
	lengthCurrentTunnerArr := len(currentTunnerArr)

	for i := 0; i < 50; i++ {

		randomNextRoomIndex := rand.Intn(lengthCurrentTunnerArr)
		nextRandomAvailableRoomName := currentTunnerArr[randomNextRoomIndex]
		// check if I visited before
		if slices.Contains(theAnt.VisitedRoomsArr, nextRandomAvailableRoomName) {
			continue
		}

		nextRoom := theGraph.Rooms[nextRandomAvailableRoomName]
		// check if the next selected room empty seat
		if nextRoom.EmptySeats == 0 {
			continue
		}

		//--------------------------------------
		/*
			if lastUsedTunnel == selectedTunnel {
			continue
			}
		*/
		// if nextRoom.LastAntCameFromRoomName==theAnt.CurrentRoomName {
		// 	fmt.Println("nextRoom.LastAntCameFromRoomName= ",nextRoom.LastAntCameFromRoomName)
		// 	continue
		// }
		//-------------------------------------
		return nextRandomAvailableRoomName

	}

	//nextRoom.LastAntCameFromRoomName =""
	// no avaliable next room
	return ""
}

//---------------------------------------------
