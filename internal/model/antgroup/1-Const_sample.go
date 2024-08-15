package antgroup

import (
	"main/internal/model/antgroup/ant"
)

type SampleAntGroup struct{}

var (
	mySampleAntGroup = SampleAntGroup{}
)

// ==================Samples=============================
func (s *SampleAntGroup) Ant_1_initmode_room_0() AntGroup {
	return AntGroup{
		NumberOfAnts: 1,
		AntsDb: map[Mtag]Ant{
			1: ant.Sample_ant1(),
		},
		CurrentSequenceNumber: 0,
		UsedTunnel:            make(map[Mtag][]Mtag),
		NotArrivedAntsName:    map[Mtag]struct{}{1: {}},
	}
}

// ----------------------
func (s *SampleAntGroup) Ant2_initmode_room_0() AntGroup {
	return AntGroup{
		NumberOfAnts: 2,
		AntsDb: map[Mtag]Ant{
			1: ant.Sample_ant1(),
			2: ant.Sample_ant2(),
		},
		CurrentSequenceNumber: 0,
		UsedTunnel:            make(map[Mtag][]Mtag),
		NotArrivedAntsName:    map[Mtag]struct{}{1: {}, 2: {}},
	}
}

//--------------------------------------
