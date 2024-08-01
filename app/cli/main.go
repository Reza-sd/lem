package main

import (
	"fmt"

	"main/internal/logstack"
	"main/internal/simulationpk/modelpk"
	"main/internal/simulationpk/randomsimulatorpk"
	"time"
)

// ----------------const----------------
const (
	pkgName = "main"
)

// ------------------logger-instance------------
var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

// --------------------------------------------
func main() {

	//------------------------
	// lem1 := simulationpk.Lem{
	// 	NumberOfAnts: 3,
	// 	StartRoom:    "1",
	// 	EndRoom:      "0",
	// 	TunnelArr:    []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5", "4-2", "2-1", "7-6", "7-2", "7-4", "6-5"},
	// }

	// lem2 := simulationpk.Lem{
	// 	NumberOfAnts: 5,
	// 	StartRoom:    "0",
	// 	EndRoom:      "1",
	// 	TunnelArr:    []string{"0-2","2-3","3-1"},
	// }
	// lem1 := modelpk.Lem{
	// 	NumberOfAnts: 5,
	// 	StartRoom:    "0",
	// 	EndRoom:      "2",
	// 	TunnelArr:    []string{"0-1", "1-2"},
	// }
	//----------------------------
	lem1 := modelpk.Lem{
		NumberOfAnts: 3,
		StartRoom:    "0",
		EndRoom:      "1",
		TunnelArr:    []string{"0-1"},
	}

	//------------------------
	/*
	   "0-2","2-3","3-1"
	*/

	//lem1=lem3
	baseModel1, errBaseModel1 := modelpk.ModelGeneratorA(&lem1)
	if errBaseModel1 != nil {
		fmt.Println("errBaseModel1 ", errBaseModel1)
		return
	}
	// -----------------------------------------
	startTimeSinceCallRandomSimulator := time.Now()
	theBestFoundTravelPlan := randomsimulatorpk.RandomSimulator(&baseModel1)
	durationRandomSimulator := time.Since(startTimeSinceCallRandomSimulator)
	//-----------------------------------
	//report
	fmt.Println("")
	fmt.Println("Steps :", theBestFoundTravelPlan.FinalSequence)
	theBestFoundTravelPlan.TheBestPlan.PrintAllAnts()
	//-----------------------------------
	fmt.Println("time : (", durationRandomSimulator.Seconds(), "Seconds =", durationRandomSimulator.Minutes(), "Minutes)")

}
