package graphpk

type SampleGragh struct{}

func (mySampleGraph *SampleGragh) Sample_graph_1() *Graph {
	// ------------------------
	//var a int = 777777
	myTunnelMap := TunnelMap{
		0: {1},
		1: {0},
	}
	//------------------------
	return &Graph{

		EndRoomName: 1,

		RoomAvailableDb: map[Mtg]bool{
			0: true,
			1: true,
		},
		TunnelsDb: &myTunnelMap,
	}

}
