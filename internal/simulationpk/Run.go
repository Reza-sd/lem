package simulationpk

import (
	//"fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
)

func (theModel *Model) Run(NumberOfTry int)(finalTravelPlan antpk.TravelPlan) {

	baseGraph := theModel.baseGraph
	baseAnts := theModel.baseAnts

	var myTravelPlan antpk.TravelPlan
	var tempGraph graphpk.Graph
	var tempAnts antpk.Ants

	

	for i := 1; i <= NumberOfTry; i++ {

		graphpk.GraphCopy(baseGraph, &tempGraph)
		antpk.AntsCopy(baseAnts, &tempAnts)

		tempAnts.TryPushAllAntsToEnd(&tempGraph, &myTravelPlan)

		tempAnts = antpk.Ants{} //reset myAnts after each try
		tempGraph = graphpk.Graph{}

	}

	// fmt.Println("")
	// fmt.Println("Steps: ", myTravelPlan.Steps)

	//myTravelPlan.SuccessfulPlan.PrintAllAnts()
	return myTravelPlan
}
