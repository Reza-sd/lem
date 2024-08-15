package antgroup

import (
	"main/internal/logstack"
	"main/internal/model/antgroup/ant"
	//"structs"
)

// ====================struct====data structure==================
type Mtag = uint16 //my type ant group pk
type Ant = ant.Ant

type AntGroup struct {
	NumberOfAnts          Mtag
	CurrentSequenceNumber Mtag
	AntsDb                map[Mtag]Ant
	//UsedTunnelinThisSeq	map[string]string//
	UsedTunnel         map[Mtag][]Mtag
	NotArrivedAntsName map[Mtag]struct{}
}
//-------------
type SampleAntGroup struct{}
//======================================================

const (
	//MaxHandleableAntsNumber = 1100
	pkgName = "antpk"
	//LogFilesDirectory =
	rootFromAntpk           = "../.."
	LogFilesDirectory       = rootFromAntpk + "/logs/"
	MaxHandleableAntsNumber = Mtag(1100)
)

var (
	mySample = SampleAntGroup{}
	
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

//======================================================
