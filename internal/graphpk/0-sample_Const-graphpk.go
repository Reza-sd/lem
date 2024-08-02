package graphpk

var (
	//------------------------
	Sample_room_0_startRoom = Room{
		Name:       "room_0",
		MaxSeats:   1000,
		EmptySeats: 1000,
		Tunnels:    []string{"room_1"},
	}
	//------------------------
	Sample_room_1_EndRoom = Room{
		Name:       "room_1",
		MaxSeats:   1000,
		EmptySeats: 1000,
		Tunnels:    []string{"room_0"},
	}
	// ------------------------
	Sample_graph_1 = Graph{
		StartRoomName:        "room_0",
		EndRoomName:          "room_1",
		CurrentAntsInEndRoom: 0,
		NumberOfAllRoom:      2,
		Rooms:                map[string]Room{
			"room_0":Sample_room_0_startRoom,
			"room_1":Sample_room_1_EndRoom,
		},
	}

//------------------------

)
