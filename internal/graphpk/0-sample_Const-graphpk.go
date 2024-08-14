package graphpk

var (
	//------------------------
	// Sample_room_0_startRoom = Room{
	// 	//Name: 0,
	// 	//MaxSeats:    1000,
	// 	HasEmptySeat: true,
	// 	//Connections:  []mtg{1},
	// }
	// //------------------------
	// Sample_room_1_EndRoom = Room{
	// 	//Name: 1,
	// 	//MaxSeats:    1000,
	// 	HasEmptySeat: true,
	// 	//Connections:  []mtg{0},
	// }
	// ------------------------
	Sample_graph_1 = Graph{
		//StartRoomName:        "room_0",
		EndRoomName:          1,
		//CurrentAntsInEndRoom: 0,
		NumberOfAllRoom:      2,
		RoomHasEmptySeatDb: map[mtg]bool{
			0: true,
			1: true,
		},
	}

//------------------------

)
