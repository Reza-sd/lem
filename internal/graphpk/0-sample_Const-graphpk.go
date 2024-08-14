package graphpk

var (

	// ------------------------
	myTunnelMap = TunnelMap{
		0: {1},
		1: {0},
	}
	//------------------------
	Sample_graph_1 = Graph{

		EndRoomName: 1,

		RoomAvailableDb: map[mtg]bool{
			0: true,
			1: true,
		},
		TunnelsDb: &myTunnelMap,
	}

//------------------------

)
