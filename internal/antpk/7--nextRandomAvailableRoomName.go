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
func (theAnt *Ant) nextRandomAvailableRoomName(theGraph *graphpk.Graph, travelHistory *TravelHistory) (string, error) {
	funcName := "nextRandomAvailableRoomName"
	//-------------------
	currentRoomObjectFromGraph, ok1 := theGraph.Rooms[theAnt.CurrentRoomName]
	if !ok1 {
		return "", logger.RWarnStr(funcName, "ok1", "the graph does not have this room", "check if graph has this room name")
	}
	//----------------
	currentTunnerArr := currentRoomObjectFromGraph.Tunnels
	lengthCurrentTunnerArr := len(currentTunnerArr)
	//------------

	for i := 0; i < 50; i++ {

		randomNextRoomIndex := rand.Intn(lengthCurrentTunnerArr)
		nextRandomAvailableRoomName := currentTunnerArr[randomNextRoomIndex]

		//---check if I visited before---------
		if slices.Contains(theAnt.VisitedRoomsArr, nextRandomAvailableRoomName) {
			continue
		}
		//---to check if the offered tunnel(from,to) is used in this sequence----
		tunnelArr, ok := travelHistory.UsedTunnels[theAnt.StepNumber+1]
		if ok {
			if tunnelArr[theAnt.CurrentRoomName] == nextRandomAvailableRoomName {
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

		//-------------------------------------
		return nextRandomAvailableRoomName, nil

	}

	return "", nil
}
