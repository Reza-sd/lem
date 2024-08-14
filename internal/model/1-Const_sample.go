package modelpk

import (
	//"fmt"
	antpk "main/internal/model/antgroup"
	graphpk "main/internal/model/graph"
)

// =====================================
func Sample_Model1() Model {

	return Model{
		Graph:    graphpk.Sample_graph_1(),
		AntGroup: antpk.Sample_AntGroup_2ant_initmode_room_0(),
	}

}

// -------------------------
func Sample_Lem_1ant_2room() Lem {
	return Lem{
		NumberOfAnts: 1,
		EndRoom:      1,
		TunnelMap: map[Mtm][]Mtm{
			0: {1},
			1: {0},
		},
	}
}

// -------------------------
func Sample_Lem_2ant_2room() Lem {
	return Lem{
		NumberOfAnts: 2,
		EndRoom:      1,
		TunnelMap: map[Mtm][]Mtm{
			0: {1},
			1: {0},
		},
	}
}

// -------------------------
func Sample_Lem_2ant_3room() Lem {
	return Lem{
		NumberOfAnts: 2,
		EndRoom:      1,
		TunnelMap: map[Mtm][]Mtm{
			0: {1},
			1: {0, 2},
			2: {1},
		},
	}
}

//-------------------------
