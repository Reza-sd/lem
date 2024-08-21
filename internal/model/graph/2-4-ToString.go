package graphpk

import (
	"fmt"
)

// ---------------------------------

func (myGraph *Graph) ToString() string {

	graphString := fmt.Sprintf("StartRoomName= 0, EndRoomName= %v \nNumberOfAllRoom= %v, \nRoomAvailableDb= %v\nTunnelsDb=%v", myGraph.EndRoomName, len(myGraph.RoomAvailableDb), myGraph.RoomAvailableDb, myGraph.TunnelsDb)

	return graphString
}

// ---------------------------------
