package antpk

import (
	_ "main/internal/graphpk"
	"main/internal/logstack"
)

// ====================struct====data structure==================
type t = uint16

type Ant struct {
	Name            t
	CurrentRoomName t
	VisitedRoomsArr []t
	StepNumber      t
}

type AntGroup struct {
	NumberOfAnts          t
	CurrentSequenceNumber t
	AntsDb                map[t]Ant
	//UsedTunnelinThisSeq	map[string]string//
	UsedTunnel         map[t][]t
	NotArrivedAntsName map[t]struct{}
}

//======================================================

const (
	MaxHandleableAntsNumber = 200
	pkgName                 = "antpk"
	//LogFilesDirectory =
	rootFromAntpk     = "../.."
	LogFilesDirectory = rootFromAntpk + "/logs/"
)

var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

//======================================================
