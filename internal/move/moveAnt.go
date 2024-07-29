package move

import (
	"fmt"
	data "main/internal/data"
	"math/rand"
	"slices"
)

// ------------------------------------------------
func (theAnt *Ant) MoveOneStepForwardRandomly(theGraph *data.Graph) {
	//fmt.Println("move theAnt")

	// -----check if its already in End room----
	if theAnt.CurrentRoomName == theGraph.EndRoomName {
		return
	}

	nextRandomAvailableRoomName := nextRandomAvailableRoomName(theAnt, theGraph)

	if nextRandomAvailableRoomName == "" {
		theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, "*")
		return
	}

	currentRoom := theGraph.Rooms[theAnt.CurrentRoomName]
	currentRoom.EmptySeats = 1
	theGraph.Rooms[theAnt.CurrentRoomName] = currentRoom
	//fmt.Println()
	nextRoom := theGraph.Rooms[nextRandomAvailableRoomName]
	nextRoom.EmptySeats = 0

	if nextRandomAvailableRoomName == theGraph.EndRoomName {
		nextRoom.EmptySeats = 1000
	}
	theGraph.Rooms[nextRandomAvailableRoomName] = nextRoom
	theAnt.VisitedRoomsArr = append(theAnt.VisitedRoomsArr, nextRandomAvailableRoomName)
	theAnt.CurrentRoomName = nextRandomAvailableRoomName
	fmt.Println(theAnt)
}

// ------------------------------------------------
func nextRandomAvailableRoomName(theAnt *Ant, theGraph *data.Graph) string {
	for i := 0; i < 50; i++ {
		currentRoom := theGraph.Rooms[theAnt.CurrentRoomName]
		currentTunnerArr := currentRoom.Tunnels
		lengthCurrentTunnerArr := len(currentTunnerArr)
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
		return nextRandomAvailableRoomName

	}
	// no avaliable next room
	return ""
}

//---------------------------------------------
