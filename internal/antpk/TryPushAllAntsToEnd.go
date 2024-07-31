package antpk

import (
	//"fmt"
	//"fmt"
	"main/internal/graphpk"
)

// =====================TryPushAllAntsToEnd===================================
func (allAnts *Ants) TryPushAllAntsToEnd(theGraph *graphpk.Graph, theTravelPlan *TravelPlan, maxTry int) {
	counter := 0
	//--------------------------------
	for {
		//if all ants have arrived
		if theGraph.CurrentAntsInEndRoom == allAnts.AntsNumber {
			break
		}

		if counter >= maxTry { // max try

			return
		}

		allAnts.MoveAllAntsOneStepRandomly(theGraph)
		counter++
	}
	//---------------------------------
	// check if steps are less than previous steps
	if counter >= theTravelPlan.Steps && theTravelPlan.Steps != 0 {
		return
	}

	theTravelPlan.Steps = counter
	theTravelPlan.TheBestPlan = *allAnts

}
