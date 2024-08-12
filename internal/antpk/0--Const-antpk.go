package antpk

import (
	"main/internal/logstack"
)

type Ant struct {
	Name            uint16
	CurrentRoomName uint16
	VisitedRoomsArr []uint16
	StepNumber      uint16
}

type AntGroup struct {
	NumberOfAnts          int
	CurrentSequenceNumber int
	AntsDb                map[string]Ant
	//UsedTunnelinThisSeq	map[string]string//
	UsedTunnel         map[string][]string
	NotArrivedAntsName map[string]struct{}
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
