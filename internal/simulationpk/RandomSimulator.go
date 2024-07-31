package simulationpk

import (
	//"fmt"
	"fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
)

// ====================================================================
func RandomSimulator(theModel *Model) (theBestFoundTravelPlan antpk.TravelPlan) {

	var myTravelPlan antpk.TravelPlan
	maxTryToPushAllAntsToEnd := theModel.baseAnts.AntsNumber * theModel.baseGraph.NumberOfAllRoom * 10
	// runRandomSimulator(theModel *Model,NumberOfTry int)
	for {

		myTravelPlan = runRandomSimulator(theModel, maxTryToPushAllAntsToEnd)
		//myTravelPlan = simulationpk.RandomSimulator(&model1)

		if myTravelPlan.Steps != 0 {
			break
		}
	}

	return myTravelPlan
}

// ====================================================================
func runRandomSimulator(theModel *Model, NumberOfTry int) (finalTravelPlan antpk.TravelPlan) {

	baseGraph := theModel.baseGraph
	baseAnts := theModel.baseAnts

	var tempTravelPlan antpk.TravelPlan
	var tempGraph graphpk.Graph
	var tempAnts antpk.Ants

	for i := 1; i <= NumberOfTry; i++ {

		graphpk.GraphCopy(baseGraph, &tempGraph)
		antpk.AntsCopy(baseAnts, &tempAnts)

		tempAnts.TryPushAllAntsToEnd(&tempGraph, &tempTravelPlan, NumberOfTry)

	}
	// ------------------
	fmt.Println("maxTryToPushAllAntsToEnd=", NumberOfTry)
	//--------------------
	return tempTravelPlan
}

//===============================================================================
