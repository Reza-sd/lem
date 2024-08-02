package randomsimulatorpk

import (
	//"fmt"
	"fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
	"main/internal/simulationpk/modelpk"
)

// ====================================================================
func RandomSimulator(theModel *modelpk.Model, travelHistory *antpk.TravelHistory) (theBestFoundTravelPlan antpk.TravelPlan) {

	var myTravelPlan antpk.TravelPlan
	maxTryToPushAllAntsToEnd := theModel.BaseAnts.NumberOfAnts * theModel.BaseGraph.NumberOfAllRoom * 10
	// runRandomSimulator(theModel *Model,NumberOfTry int)
	for {

		myTravelPlan = runRandomSimulator(theModel, maxTryToPushAllAntsToEnd, travelHistory)
		//myTravelPlan = simulationpk.RandomSimulator(&model1)

		if myTravelPlan.FinalSequence != 0 {
			break
		}
	}

	return myTravelPlan
}

// ====================================================================
func runRandomSimulator(theModel *modelpk.Model, NumberOfTry int, travelHistory *antpk.TravelHistory) (finalTravelPlan antpk.TravelPlan) {

	baseGraph := theModel.BaseGraph
	baseAnts := theModel.BaseAnts

	var tempTravelPlan antpk.TravelPlan
	var tempGraph graphpk.Graph
	var tempAnts antpk.AntGroup

	for i := 1; i <= NumberOfTry; i++ {

		graphpk.GraphCopy(baseGraph, &tempGraph)
		antpk.AntGroupCopy(baseAnts, &tempAnts)

		tempAnts.TryPushAllAntsToEnd(&tempGraph, &tempTravelPlan, NumberOfTry, travelHistory)

	}
	// ------------------
	fmt.Println("maxTryToPushAllAntsToEnd=", NumberOfTry)
	//--------------------
	return tempTravelPlan
}

//===============================================================================
