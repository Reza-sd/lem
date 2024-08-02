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
	NumberOfAnts     int
	AntsMap          map[int]Ant
	NumberOfSequence int
}

type TravelPlan struct {
	FinalSequence int
	TheBestPlan   AntGroup
}

type TravelHistory struct {
	UsedTunnels map[int]map[string]string
	//seq 3 roomA:roomB
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
