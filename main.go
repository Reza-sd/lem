package main

import (
	"fmt"
	"main/internal/simulationpk"
	"time"
)

func main() {

	lem1 := simulationpk.Lem{
		NumberOfAnts: 10,
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

	/*
	   "0-2","2-3","3-1"
	*/
	startTime := time.Now() //
	//lem1=lem2
	model1 := lem1.ModelInit()
	myTravelPlan := model1.Run(20000)
	fmt.Println("")
	fmt.Println("Steps :", myTravelPlan.Steps)
	myTravelPlan.TheBestPlan.PrintAllAnts()
	fmt.Println("time =", time.Since(startTime).Seconds(), "Seconds")
	//fmt.Println("time=",time.Since(startTime).Minutes(),"Minutes")
	//fmt.Println("time=",time.Since(startTime).Abs().Hours(),"Abs Hours")
}
