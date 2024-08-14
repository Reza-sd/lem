package graphpk

// =======================================================

func (myGraph *Graph) InitGraph(tunnelMap map[mtg][]mtg, EndRoomName mtg) error {

	myGraph.EndRoomName = EndRoomName

	myGraph.RoomHasEmptySeatDb = make(map[mtg]bool)
	myGraph.TunnelsDb = make(map[mtg][]mtg)

	myGraph.TunnelsDb = tunnelMap

	for key := range myGraph.TunnelsDb {
		myGraph.RoomHasEmptySeatDb[key] = true
	}

	return nil
}

//=======================================================
