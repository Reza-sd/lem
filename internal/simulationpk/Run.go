package simulationpk

import (
	"fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
)


func (myLem *Lem)Run(NumberOfTry int){


	NumberOfAnts := myLem.NumberOfAnts
	StartRoom :=myLem.StartRoom
	EndRoom :=myLem.EndRoom

	
	tunnelArr := myLem.TunnelArr

	var baseGraph graphpk.Graph
	baseGraph.InitGraph(tunnelArr, StartRoom, EndRoom)

	var baseAnts antpk.Ants
	baseAnts.AntsInit(NumberOfAnts, &baseGraph)


	var myTravelPlan antpk.TravelPlan
	var myGraph graphpk.Graph
	var myAnts antpk.Ants

	myTravelPlan = antpk.TravelPlan{Steps: 10000}

	for i := 1; i <= NumberOfTry; i++ {
		
		graphpk.GraphCopy(baseGraph,&myGraph)
		antpk.AntsCopy(baseAnts, &myAnts)

		myAnts.TryPushAllAntsToEnd(&myGraph, &myTravelPlan)

		myAnts = antpk.Ants{} //reset myAnts after each try
		myGraph = graphpk.Graph{}
	

	}
	
	fmt.Println("")
	fmt.Println("Steps: ", myTravelPlan.Steps)
	
	myTravelPlan.SuccessfulPlan.PrintAllAnts()
}