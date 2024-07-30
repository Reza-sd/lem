package main

import (
	"fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
)

func main() {

	var baseGraph graphpk.Graph
	tunnelArr := []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5", "4-2", "2-1", "7-6", "7-2", "7-4", "6-5"}
	baseGraph.InitGraph(tunnelArr, "1", "0")

	var baseAnts antpk.Ants
	baseAnts.AntsInit(3, &baseGraph)


	var myTravelPlan antpk.TravelPlan
	var myGraph graphpk.Graph
	var myAnts antpk.Ants

	myTravelPlan = antpk.TravelPlan{Steps: 10000}

	for i := 1; i <= 200; i++ {
		//fmt.Println("try=", i)
		//myTravelPlan = antpk.TravelPlan{Steps: 10000}

		
	
		graphpk.GraphCopy(baseGraph,&myGraph)
		antpk.AntsCopy(baseAnts, &myAnts)

		//myGraph.PrintGraph()
	
		//myAnts.PrintAllAnts()
		

		myAnts.TryPushAllAntsToEnd(&myGraph, &myTravelPlan)

		//fmt.Println("number Steps: ", myTravelPlan.Steps)
		//myTravelPlan.SuccessfulPlan.PrintAllAnts()

		myAnts = antpk.Ants{} //reset myAnts after each try
		myGraph = graphpk.Graph{}
	

	}
	fmt.Println("")
	fmt.Println("Steps: ", myTravelPlan.Steps)
	
	myTravelPlan.SuccessfulPlan.PrintAllAnts()

}
