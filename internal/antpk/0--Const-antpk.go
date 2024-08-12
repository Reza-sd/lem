package antpk

import (
	"main/internal/logstack"
)

type t=uint16

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

// type TravelHistory struct {
// 	UsedTunnelsMap map[int]map[string]string
// }

// type TravelPlan struct {
// 	FinalSequence int
// 	TheBestPlan   AntGroup
// }

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
