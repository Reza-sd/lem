package antpk

import (
	_ "main/internal/graphpk"
	"main/internal/logstack"
)

// ====================struct====data structure==================
type mta = uint16 //my type ant pk

type Ant struct {
	//Name            mta
	CurrentRoomName mta
	VisitedRoomsArr []mta
	StepNumber      mta
}

type AntGroup struct {
	NumberOfAnts          mta
	CurrentSequenceNumber mta
	AntsDb                map[mta]Ant
	//UsedTunnelinThisSeq	map[string]string//
	UsedTunnel         map[mta][]mta
	NotArrivedAntsName map[mta]struct{}
}

//======================================================

const (
	//MaxHandleableAntsNumber = 1100
	pkgName = "antpk"
	//LogFilesDirectory =
	rootFromAntpk     = "../.."
	LogFilesDirectory = rootFromAntpk + "/logs/"
)

var (
	MaxHandleableAntsNumber = mta(1100)
	logger                  = logstack.LogCollector{
		PackageName: pkgName,
	}
)

//======================================================
