package modelpk

import (
	antgrouppk "main/internal/model/antgroup"
	graphpk "main/internal/model/graph"
)

type SampleModel struct{}

var (
	mySampleGraph    = graphpk.SampleGragh{}
	mySampleAntGroup = antgrouppk.SampleAntGroup{}

	mySampleModel = SampleModel{}
)

// =====================================
func (s *SampleModel) Sample_Model1() Model {

	return Model{
		Graph:    mySampleGraph.Sample_graph_1(),
		AntGroup: mySampleAntGroup.Ant2_initmode_room_0(),
	}

}

// -------------------------
func (s *SampleModel) Sample_Lem_1ant_2room() Lem {
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
func (s *SampleModel) Sample_Lem_2ant_2room() Lem {
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
func (s *SampleModel) Sample_Lem_2ant_3room() Lem {
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
