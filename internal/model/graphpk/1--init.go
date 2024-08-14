package graphpk

// =======================================================

func (myGraph *Graph) Init(myTunnelMap TunnelMap, EndRoomName Mtg) error {

	myGraph.EndRoomName = EndRoomName

	myGraph.RoomAvailableDb = make(map[Mtg]bool)

	myGraph.TunnelsDb = &myTunnelMap

	for key := range *myGraph.TunnelsDb {
		myGraph.RoomAvailableDb[key] = true
	}

	return nil
}

//=======================================================
