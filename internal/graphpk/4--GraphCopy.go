package graphpk

func GraphCopyFresh(baseGraph Graph, secondGraph *Graph) {

	secondGraph.StartRoomName = baseGraph.StartRoomName
	secondGraph.EndRoomName = baseGraph.EndRoomName
	secondGraph.CurrentAntsInEndRoom = 0
	secondGraph.NumberOfAllRoom = baseGraph.NumberOfAllRoom
	//----------------------------------
	secondGraph.Rooms = make(map[string]Room)
	var room Room

	for roomName, value := range baseGraph.Rooms {

		room.Name = value.Name
		room.MaxSeats = value.MaxSeats
		room.EmptySeats = value.EmptySeats

		//copy(value.Tunnels,room.Tunnels)
		room.Tunnels = deepCopySlice(value.Tunnels)
		secondGraph.Rooms[roomName] = room

	}
}

// ======================
func deepCopySlice(original []string) []string {
	copy := make([]string, len(original))
	for i, item := range original {
		copy[i] = item
	}
	return copy
}

//======================
