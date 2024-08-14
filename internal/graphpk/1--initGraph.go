package graphpk

// "fmt"

// =======================================================
func (myGraph *Graph) InitGraph(tunnelMap map[mtg][]mtg, EndRoomName mtg) error {
	//funcName := "InitGraph"
	// set start and end room
	//myGraph.StartRoomName = start
	myGraph.EndRoomName = EndRoomName
	//myGraph.CurrentAntsInEndRoom = 0
	// Initialize rooms map
	myGraph.RoomHasEmptySeatDb = make(map[mtg]bool)
	myGraph.TunnelsDb = make(map[mtg][]mtg)

	myGraph.TunnelsDb = tunnelMap

	for key := range myGraph.TunnelsDb {
		myGraph.RoomHasEmptySeatDb[key] = true
	}
	//err := myGraph.setTunnelsToGraphRooms(tunnelMap)
	// if err != nil {
	// 	//fmt.Println("Error:", err)
	// 	return logger.RWarn(funcName, "setTunnelAndSeatsToGraphRooms", err, "setTunnelAndSeatsToGraphRooms ")
	// }
	myGraph.NumberOfAllRoom = mtg(len(myGraph.RoomHasEmptySeatDb))
	return nil
}

//=======================================================
