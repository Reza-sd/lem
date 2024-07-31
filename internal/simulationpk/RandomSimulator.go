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
	// runRandomSimulator(theModel *Model,NumberOfTry int)
	for {

		myTravelPlan = runRandomSimulator(theModel, 200)
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
	//maxTry := 200 //?
	maxTryToPushAllAntsToEnd := theModel.baseAnts.AntsNumber * theModel.baseGraph.NumberOfAllRoom * 10

	var tempTravelPlan antpk.TravelPlan
	var tempGraph graphpk.Graph
	var tempAnts antpk.Ants

	for i := 1; i <= NumberOfTry; i++ {

		graphpk.GraphCopy(baseGraph, &tempGraph)
		antpk.AntsCopy(baseAnts, &tempAnts)

		tempAnts.TryPushAllAntsToEnd(&tempGraph, &tempTravelPlan, maxTryToPushAllAntsToEnd)

	}
	// ------------------
	fmt.Println("maxTryToPushAllAntsToEnd=", maxTryToPushAllAntsToEnd)
	//--------------------
	return tempTravelPlan
}

//===============================================================================
