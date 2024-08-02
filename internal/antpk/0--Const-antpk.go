package antpk

import (
	"main/internal/logstack"
)

type Ant struct {
	Name            string
	CurrentRoomName string
	VisitedRoomsArr []string
	StepNumber      int
	//AssignedPath    []string
}

type AntGroup struct {
	NumberOfAnts   int
	SequenceNumber int
	AntsMap        map[int]Ant
	UsedTunnel     TravelHistory
}

type TravelPlan struct {
	FinalSequence int
	TheBestPlan   AntGroup
}

type TravelHistory struct {
	UsedTunnelsMap map[int]map[string]string
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
