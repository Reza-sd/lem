package graphpk

func GraphCopyFresh(original Graph) Graph {

	// secondGraph.StartRoomName = baseGraph.StartRoomName
	// secondGraph.EndRoomName = baseGraph.EndRoomName
	// secondGraph.CurrentAntsInEndRoom = 0
	// secondGraph.NumberOfAllRoom = baseGraph.NumberOfAllRoom
	// //----------------------------------
	// secondGraph.Rooms = make(map[string]Room)
	
	copy := Graph{
        StartRoomName:        original.StartRoomName,
        EndRoomName:          original.EndRoomName,
        CurrentAntsInEndRoom: original.CurrentAntsInEndRoom,
        NumberOfAllRoom:      original.NumberOfAllRoom,
        Rooms:                make(map[string]Room, len(original.Rooms)),
    }
	//var room Room

	for name, room := range original.Rooms {

		copy.Rooms[name] = Room{
			Name:       room.Name,
			MaxSeats:   room.MaxSeats,
			EmptySeats: room.EmptySeats,
			Tunnels:    append([]string(nil), room.Tunnels...),
		}

		// room.Name = value.Name
		// room.MaxSeats = value.MaxSeats
		// room.EmptySeats = value.EmptySeats

		//copy(value.Tunnels,room.Tunnels)
		// room.Tunnels= make([]string, len(value.Tunnels))
		// copy(value.Tunnels,room.Tunnels)
		// room.Tunnels = deepCopySlice(value.Tunnels)
		// secondGraph.Rooms[roomName] = room

	}
	return copy
}

// ======================
// func deepCopySlice(original []string) []string {
// 	copy := make([]string, len(original))
// 	for i, item := range original {
// 		copy[i] = item
// 	}
// 	return copy
// }

//======================
