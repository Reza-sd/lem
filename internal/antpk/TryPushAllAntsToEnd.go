package antpk

import (
	//"fmt"
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *Ants) TryPushAllAntsToEnd(theGraph *graphpk.Graph, theTravelPlan *TravelPlan) {
	counter := 0
	//success := true
	for {

		if theGraph.CurrentAntsInEndRoom == allAnts.AntsNumber {
			break
		}

		if counter >= 10 { // max try
		
			return
		}

		allAnts.MoveAllAntsOneStepRandomly(theGraph)
		counter++
	}
	
	if counter >= theTravelPlan.Steps && theTravelPlan.Steps  != 0 {
		return
	}

		theTravelPlan.Steps = counter
		theTravelPlan.SuccessfulPlan = *allAnts


}
