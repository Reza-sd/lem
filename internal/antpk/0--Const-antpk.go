package antpk

import (
	_ "main/internal/graphpk"
	"main/internal/logstack"
)

// ====================struct====data structure==================
type mt = uint16 //my type

type Ant struct {
	Name            mt
	CurrentRoomName mt
	VisitedRoomsArr []mt
	StepNumber      mt
}

type AntGroup struct {
	NumberOfAnts          mt
	CurrentSequenceNumber mt
	AntsDb                map[mt]Ant
	//UsedTunnelinThisSeq	map[string]string//
	UsedTunnel         map[mt][]mt
	NotArrivedAntsName map[mt]struct{}
}

//======================================================

const (
	//MaxHandleableAntsNumber = 1100
	pkgName                 = "antpk"
	//LogFilesDirectory =
	rootFromAntpk     = "../.."
	LogFilesDirectory = rootFromAntpk + "/logs/"
)

var (
	MaxHandleableAntsNumber = mt(1100)
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

//======================================================
