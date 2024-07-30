package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *Ants) PushAllAntsToEnd(theGraph *graphpk.Graph, theTravelPlan *TravelPlan) {
	counter := 0
	success := true
	for {
		counter++
		if counter>=10{ // max try
			success = false
			break
		}
		if theGraph.CurrentAntsInEndRoom == allAnts.AntsNumber {
			break
		}
		allAnts.MoveAllAntsOneStepRandomly(theGraph)

	}
	//allAnts
	if !success { return}

	theTravelPlan.Steps = counter
	theTravelPlan.SuccessfulPlan = *allAnts
}
