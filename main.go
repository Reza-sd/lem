package main

import (
	"fmt"
	"main/internal/simulationpk"
)



func main() {

	lem1 := simulationpk.Lem{
		NumberOfAnts: 3,
		StartRoom:    "1",
		EndRoom:      "0",
		TunnelArr:    []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5", "4-2", "2-1", "7-6", "7-2", "7-4", "6-5"},
	}
	model1:= lem1.ModelInit()
	myTravelPlan := model1.Run(200)
	fmt.Println("")
	fmt.Println("Steps: ", myTravelPlan.Steps)
	myTravelPlan.SuccessfulPlan.PrintAllAnts()

}
