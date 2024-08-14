package modelpk

import (
	//"fmt"
	antpk "main/internal/model/antpk"
	graphpk "main/internal/model/graphpk"
)

func Sample_Model1() Model {

	return Model{
		Graph:    graphpk.Sample_graph_1(),
		AntGroup: antpk.Sample_AntGroup_2ant_initmode_room_0(),
	}

}
func Sample_Lem1() Lem {
	return Lem{
		NumberOfAnts: 1,
		EndRoom:      1,
		TunnelMap: map[Mtm][]Mtm{
			0: {1},
			1: {0},
		},
	}
}

//-------------------------
