package graphpk

import (
	"fmt"
)

// ---------------------------------

func (myGraph *Graph) ToString() string {

	graphString := fmt.Sprintf("StartRoomName= 0, EndRoomName= %v \nNumberOfAllRoom= %v, \nRoomHasEmptySeatDb= %v\nTunnelsDb=%v", myGraph.EndRoomName, len(myGraph.RoomHasEmptySeatDb), myGraph.RoomHasEmptySeatDb, myGraph.TunnelsDb)

	return graphString
}

// ---------------------------------
