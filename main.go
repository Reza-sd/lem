package main

import (
	"fmt"
	//"main/internal/antpk"
	"main/internal/simulationpk"
	"time"
)

func main() {

	lem1 := simulationpk.Lem{
		NumberOfAnts: 100,
		StartRoom:    "1",
		EndRoom:      "0",
		TunnelArr:    []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5", "4-2", "2-1", "7-6", "7-2", "7-4", "6-5"},
	}

	// lem2 := simulationpk.Lem{
	// 	NumberOfAnts: 5,
	// 	StartRoom:    "0",
	// 	EndRoom:      "1",
	// 	TunnelArr:    []string{"0-2","2-3","3-1"},
	// }
	// lem1 := simulationpk.Lem{
	// 	NumberOfAnts: 3,
	// 	StartRoom:    "0",
	// 	EndRoom:      "END",
	// 	TunnelArr:    []string{"0-3","3-END"},
	// }

	/*
	   "0-2","2-3","3-1"
	*/

	//lem1=lem3
	model1 := lem1.ModelInit()
	// -----------------------------------------
	startTimeSinceCallRandomSimulator := time.Now()
	theBestFoundTravelPlan := simulationpk.RandomSimulator(&model1)
	durationRandomSimulator := time.Since(startTimeSinceCallRandomSimulator)
	//-----------------------------------
	//report
	fmt.Println("")
	fmt.Println("Steps :", theBestFoundTravelPlan.Steps)
	theBestFoundTravelPlan.TheBestPlan.PrintAllAnts()
	//-----------------------------------
	fmt.Println("time : (", durationRandomSimulator.Seconds(), "Seconds =", durationRandomSimulator.Minutes(), "Minutes)")

}
