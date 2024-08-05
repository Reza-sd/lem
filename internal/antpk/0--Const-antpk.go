package antpk

import (
	"main/internal/logstack"
)

type Ant struct {
	Name            string
	CurrentRoomName string
	VisitedRoomsArr []string
	StepNumber      int
}

type AntGroup struct {
	NumberOfAnts          int
	currentSequenceNumber int
	AntsDb                map[string]Ant
	UsedTunnel            map[int]map[string]string
	AntsInEndRoom         []string
}

// type TravelHistory struct {
// 	UsedTunnelsMap map[int]map[string]string
// }

type TravelPlan struct {
	FinalSequence int
	TheBestPlan   AntGroup
}

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
