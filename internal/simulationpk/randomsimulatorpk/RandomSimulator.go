package randomsimulatorpk

import (
	//"fmt"
	"fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
	"main/internal/simulationpk/modelpk"
)

// ====================================================================
func RandomSimulator(theModel *modelpk.Model) (theBestFoundTravelPlan antpk.TravelPlan) {

	var myTravelPlan antpk.TravelPlan
	maxTryToPushAllAntsToEnd := theModel.BaseAnts.NumberOfAnts * theModel.BaseGraph.NumberOfAllRoom * 10
	// runRandomSimulator(theModel *Model,NumberOfTry int)
	for {

		myTravelPlan = runRandomSimulator(theModel, maxTryToPushAllAntsToEnd)
		//myTravelPlan = simulationpk.RandomSimulator(&model1)

		if myTravelPlan.FinalSequence != 0 {
			break
		}
	}

	return myTravelPlan
}

// ====================================================================
func runRandomSimulator(theModel *modelpk.Model, NumberOfTry int) (finalTravelPlan antpk.TravelPlan) {

	baseGraph := theModel.BaseGraph
	baseAnts := theModel.BaseAnts

	var tempTravelPlan antpk.TravelPlan
	var tempGraph graphpk.Graph
	var tempAnts antpk.AntGroup

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
