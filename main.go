package main

import (
	"fmt"
	//"main/internal/antpk"
	"main/internal/simulationpk/modelpk"
	"main/internal/simulationpk/randomsimulator"
	"time"
)

func main() {

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
	lem1 := modelpk.Lem{
		NumberOfAnts: 2,
		StartRoom:    "0",
		EndRoom:      "3",
		TunnelArr:    []string{"0-3"},
	}

	/*
	   "0-2","2-3","3-1"
	*/

	//lem1=lem3
	baseModel1 := modelpk.ModelGeneratorA(lem1)
	// -----------------------------------------
	startTimeSinceCallRandomSimulator := time.Now()
	theBestFoundTravelPlan := randomsimulator.RandomSimulator(&baseModel1)
	durationRandomSimulator := time.Since(startTimeSinceCallRandomSimulator)
	//-----------------------------------
	//report
	fmt.Println("")
	fmt.Println("Steps :", theBestFoundTravelPlan.Steps)
	theBestFoundTravelPlan.TheBestPlan.PrintAllAnts()
	//-----------------------------------
	fmt.Println("time : (", durationRandomSimulator.Seconds(), "Seconds =", durationRandomSimulator.Minutes(), "Minutes)")

}
