package graphpk

func GraphFreshCopy(original Graph) Graph {

	copy := Graph{
		StartRoomName:        original.StartRoomName,
		EndRoomName:          original.EndRoomName,
		CurrentAntsInEndRoom: original.CurrentAntsInEndRoom,
		NumberOfAllRoom:      original.NumberOfAllRoom,
		Rooms:                make(map[string]Room, len(original.Rooms)),
	}

	for name, room := range original.Rooms {

		copy.Rooms[name] = Room{
			Name:        room.Name,
			MaxSeats:    room.MaxSeats,
			EmptySeats:  room.EmptySeats,
			Connections: append([]string(nil), room.Connections...),
		}
	}
	return copy
}
