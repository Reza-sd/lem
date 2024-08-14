package antpk

import (
	"main/internal/logstack"
)

// ====================struct====data structure==================
type Mta = uint16 //my type ant pk

type Ant struct {
	//Name            mta
	CurrentRoomName Mta
	VisitedRoomsArr []Mta
	StepNumber      Mta
}

type AntGroup struct {
	NumberOfAnts          Mta
	CurrentSequenceNumber Mta
	AntsDb                map[Mta]Ant
	//UsedTunnelinThisSeq	map[string]string//
	UsedTunnel         map[Mta][]Mta
	NotArrivedAntsName map[Mta]struct{}
}

//======================================================

const (
	//MaxHandleableAntsNumber = 1100
	pkgName = "antpk"
	//LogFilesDirectory =
	rootFromAntpk           = "../.."
	LogFilesDirectory       = rootFromAntpk + "/logs/"
	MaxHandleableAntsNumber = Mta(1100)
)

var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

//======================================================
