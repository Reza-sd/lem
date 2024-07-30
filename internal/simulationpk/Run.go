package simulationpk

import (
	"fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
)

func (myLem *Lem) Run(NumberOfTry int) {

	baseGraph := myLem.baseModel.baseGraph
	baseAnts := myLem.baseModel.baseAnts

	var myTravelPlan antpk.TravelPlan
	var myGraph graphpk.Graph
	var myAnts antpk.Ants

	myTravelPlan = antpk.TravelPlan{Steps: 10000}

	for i := 1; i <= NumberOfTry; i++ {

		graphpk.GraphCopy(baseGraph, &myGraph)
		antpk.AntsCopy(baseAnts, &myAnts)

		myAnts.TryPushAllAntsToEnd(&myGraph, &myTravelPlan)

		myAnts = antpk.Ants{} //reset myAnts after each try
		myGraph = graphpk.Graph{}

	}

	fmt.Println("")
	fmt.Println("Steps: ", myTravelPlan.Steps)

	myTravelPlan.SuccessfulPlan.PrintAllAnts()
}
