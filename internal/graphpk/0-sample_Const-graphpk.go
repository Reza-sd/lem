package graphpk

var (

	// ------------------------
	Sample_graph_1 = Graph{
	
		EndRoomName:          1,
	
		RoomHasEmptySeatDb: map[mtg]bool{
			0: true,
			1: true,
		},
		TunnelsDb: map[mtg][]mtg{
			0:{1},
			1:{0},
		},
	}

//------------------------

)
