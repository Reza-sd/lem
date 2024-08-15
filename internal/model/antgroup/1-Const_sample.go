package antgroup

import (
	"main/internal/model/antgroup/ant"
)

type SampleAntGroup struct{}

var (
	mySampleAntGroup = SampleAntGroup{}
)

// ==================Samples=============================
func (s *SampleAntGroup) Sample_AntGroup_1ant_initmode_room_0() AntGroup {
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
func (s *SampleAntGroup) Sample_AntGroup_2ant_initmode_room_0() AntGroup {
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
