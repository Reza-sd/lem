package graphpk

// =======================================================

func (myGraph *Graph) InitGraph(myTunnelMap TunnelMap, EndRoomName mtg) error {

	myGraph.EndRoomName = EndRoomName

	myGraph.RoomAvailableDb = make(map[mtg]bool)

	myGraph.TunnelsDb = &myTunnelMap

	for key := range *myGraph.TunnelsDb {
		myGraph.RoomAvailableDb[key] = true
	}

	return nil
}

//=======================================================
