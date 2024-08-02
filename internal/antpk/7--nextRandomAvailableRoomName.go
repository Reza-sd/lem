package antpk

import (
	//"fmt"
	//"fmt"
	//"fmt"
	graphpk "main/internal/graphpk"
	"math/rand"
	"slices"
)

// ==========================================================
func (theAnt *Ant) nextRandomAvailableRoomName(theGraph *graphpk.Graph) string {
	currentRoomObjectFromGraph := theGraph.Rooms[theAnt.CurrentRoomName]
	currentTunnerArr := currentRoomObjectFromGraph.Tunnels
	lengthCurrentTunnerArr := len(currentTunnerArr)

	for i := 0; i < 50; i++ {

		randomNextRoomIndex := rand.Intn(lengthCurrentTunnerArr)
		nextRandomAvailableRoomName := currentTunnerArr[randomNextRoomIndex]

		// check if I visited before
		if slices.Contains(theAnt.VisitedRoomsArr, nextRandomAvailableRoomName) {
			continue
		}
		//---------------------

		//---------------------
		nextRoom := theGraph.Rooms[nextRandomAvailableRoomName]
		// check if the next selected room empty seat
		if nextRoom.EmptySeats == 0 {
			continue
		}
		//-----------------------------

		//-------------------------------------
		return nextRandomAvailableRoomName

	}

	return ""
}
