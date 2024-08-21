package antgroup

import (
	antpk "main/internal/model/ant"
)

type SampleAntGroup struct{}

var (
	MySampleAntGroup = SampleAntGroup{}
	MySampleAnt      = antpk.SampleAnt{}
)

// ==================Samples=============================
func (s *SampleAntGroup) Ant_1_initmode_room_0() AntGroup {
	return AntGroup{
		NumberOfAnts: 1,
		AntsDb: map[Mtag]*Ant{
			1: MySampleAnt.Ant1(),
		},
		CurrentSequenceNumber: 0,
		UsedTunnel:            make(map[Mtag][]Mtag),
		NotArrivedAntsName:    map[Mtag]struct{}{1: {}},
	}
}

// ----------------------
func (s *SampleAntGroup) Ant_2_initmode_room_0() AntGroup {
	return AntGroup{
		NumberOfAnts: 2,
		AntsDb: map[Mtag]*Ant{
			1: MySampleAnt.Ant1(),
			2: MySampleAnt.Ant2(),
		},
		CurrentSequenceNumber: 0,
		UsedTunnel:            make(map[Mtag][]Mtag),
		NotArrivedAntsName:    map[Mtag]struct{}{1: {}, 2: {}},
	}
}

//--------------------------------------
