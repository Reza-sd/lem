package simulationpk

import (
	//"fmt"
	//"fmt"
	"main/internal/antpk"
	//graphpk "main/internal/graphpk"
)

func RandomSimulator(theModel *Model) (theBestFoundTravelPlan antpk.TravelPlan) {

	var myTravelPlan antpk.TravelPlan
	// runRandomSimulator(theModel *Model,NumberOfTry int)
	for {

		myTravelPlan = runRandomSimulator(theModel, 500)
		//myTravelPlan = simulationpk.RandomSimulator(&model1)

		if myTravelPlan.Steps != 0 {
			break
		}
	}

	return myTravelPlan
}
