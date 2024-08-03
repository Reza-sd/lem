package graphpk

import (
	//"fmt"
)

// =======================================================
func (myGraph *Graph) InitGraph(tunnelArr []string, start, end string) error {
	funcName:="InitGraph"
	// set start and end room
	myGraph.StartRoomName = start
	myGraph.EndRoomName = end
	myGraph.CurrentAntsInEndRoom = 0
	// Initialize rooms map
	myGraph.Rooms = make(map[string]Room)

	err := myGraph.setTunnelAndSeatsToGraphRooms(tunnelArr)
	if err != nil {
		//fmt.Println("Error:", err)
		return logger.RWarn(funcName,"setTunnelAndSeatsToGraphRooms",err,"setTunnelAndSeatsToGraphRooms ")
	}
	myGraph.NumberOfAllRoom = len(myGraph.Rooms)
	return nil
}

//=======================================================
