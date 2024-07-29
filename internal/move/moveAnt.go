package move

import (
	"fmt"
	data "main/internal/data"
	"math/rand"
	"slices"
)

// ------------------------------------------------
func (theAnt *Ant) MoveOneStepForwardRandomly(theGraph *data.Graph) {
	fmt.Println("move theAnt")

	// -----check if its already in End room----
	if theAnt.CurrentRoomName == theGraph.EndRoomName {
		return
	}
	
	nextRandomAvailableRoomName := nextRandomAvailableRoomName(theAnt, theGraph)

	

	if nextRandomAvailableRoomName == "" {
		theAnt.VisitedRoomsArr=append(theAnt.VisitedRoomsArr, "*")
		return
	}
	
	currentRoom := theGraph.Rooms[theAnt.CurrentRoomName]
	currentRoom.EmptySeats++
	nextRoom := theGraph.Rooms[nextRandomAvailableRoomName]
	nextRoom.EmptySeats--

	theAnt.VisitedRoomsArr=append(theAnt.VisitedRoomsArr, nextRandomAvailableRoomName)
	theAnt.CurrentRoomName = nextRandomAvailableRoomName


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
