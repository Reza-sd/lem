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
