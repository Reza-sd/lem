package graphpk

func GraphCopy(baseGraph Graph, secondGraph *Graph) {

	secondGraph.StartRoomName = baseGraph.StartRoomName
	secondGraph.EndRoomName = baseGraph.EndRoomName
	secondGraph.CurrentAntsInEndRoom = 0

	//----------------------------------
	secondGraph.Rooms = make(map[string]Room)
	var room Room

	for roomName, value := range baseGraph.Rooms {

		//room = secondGraph.Rooms[roomName]
		room.name = value.name
		room.maxSeats = value.maxSeats
		room.EmptySeats = value.EmptySeats
		room.Tunnels = value.Tunnels
		//copy(room.Tunnels,value.Tunnels)

		secondGraph.Rooms[roomName] = room
		//secondGraph.Rooms[roomName] = value
	}
}
