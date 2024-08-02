package graphpk

func (myGraph *Graph) setTunnelAndSeatsToGraphRooms(tunnelArr []string) error {
	funcName := "setTunnelToGraph"
	//-----------------------------
	for _, tunnel := range tunnelArr {
		//-----------------------------
		fromRoom, toRoom, errSplitstr := splitStringToParts(tunnel)
		if errSplitstr != nil {

			return logger.RWarn(funcName, "splitStringToParts", errSplitstr, "split tunnel string to its parts")
		}
		//-----------------------------
		// Retrieve the Room from the map
		roomA := myGraph.Rooms[fromRoom]
		roomB := myGraph.Rooms[toRoom]

		roomA.Name = fromRoom
		roomB.Name = toRoom
		// Modify the Room's tunnels
		roomA.Tunnels = append(roomA.Tunnels, toRoom)
		roomB.Tunnels = append(roomB.Tunnels, fromRoom)
		//---------------set-seats----------------------
		// Put the modified Room back into the map
		if fromRoom == myGraph.StartRoomName || fromRoom == myGraph.EndRoomName {
			roomA.MaxSeats = 100000

		} else {
			roomA.MaxSeats = 1

		}

		if toRoom == myGraph.StartRoomName || toRoom == myGraph.EndRoomName {
			roomB.MaxSeats = 100000

		} else {
			roomB.MaxSeats = 1

		}
		//--------------------------------------------
		roomA.EmptySeats = roomA.MaxSeats
		roomB.EmptySeats = roomB.MaxSeats

		myGraph.Rooms[fromRoom] = roomA
		myGraph.Rooms[toRoom] = roomB

	}

	return nil
}

//--------------------------------
