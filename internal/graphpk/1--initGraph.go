package graphpk

// "fmt"

// =======================================================
func (myGraph *Graph) InitGraph(tunnelMap map[mtg][]mtg, end mtg) error {
	//funcName := "InitGraph"
	// set start and end room
	//myGraph.StartRoomName = start
	myGraph.EndRoomName = end
	myGraph.CurrentAntsInEndRoom = 0
	// Initialize rooms map
	myGraph.RoomsDb = make(map[mtg]bool)

	//err := myGraph.setTunnelsToGraphRooms(tunnelMap)
	// if err != nil {
	// 	//fmt.Println("Error:", err)
	// 	return logger.RWarn(funcName, "setTunnelAndSeatsToGraphRooms", err, "setTunnelAndSeatsToGraphRooms ")
	// }
	myGraph.NumberOfAllRoom = mtg(len(myGraph.RoomsDb))
	return nil
}

//=======================================================
