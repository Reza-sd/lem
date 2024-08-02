package graphpk

func GraphCopy(baseGraph Graph, secondGraph *Graph) {

	secondGraph.StartRoomName = baseGraph.StartRoomName
	secondGraph.EndRoomName = baseGraph.EndRoomName
	secondGraph.CurrentAntsInEndRoom = 0
	secondGraph.NumberOfAllRoom = baseGraph.CurrentAntsInEndRoom
	//----------------------------------
	secondGraph.Rooms = make(map[string]Room)
	var room Room

	for roomName, value := range baseGraph.Rooms {

		room.Name = value.Name
		room.MaxSeats = value.MaxSeats
		room.EmptySeats = value.EmptySeats
		room.Tunnels = value.Tunnels

		secondGraph.Rooms[roomName] = room

	}
}
