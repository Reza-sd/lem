package antgroup

import (
	"main/internal/logstack"
	antpk "main/internal/model/antgroup/ant"
	//"structs"
)

// ====================struct====data structure==================
type Mtag = uint16 //my type ant group pk
type Ant = antpk.Ant

type AntGroup struct {
	NumberOfAnts          Mtag
	CurrentSequenceNumber Mtag
	AntsDb                map[Mtag]Ant
	//UsedTunnelinThisSeq	map[string]string//
	UsedTunnel         map[Mtag][]Mtag
	NotArrivedAntsName map[Mtag]struct{}
}
//-------------
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

	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

//======================================================
