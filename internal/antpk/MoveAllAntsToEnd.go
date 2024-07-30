package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *Ants) PushAllAntsToEnd(theGraph *graphpk.Graph, theTravelPlan *TravelPlan) {
	counter := 0
	for {
		counter++
		if theGraph.CurrentAntsInEndRoom == allAnts.AntsNumber || counter == 10 {
			break
		}
		allAnts.MoveAllAntsOneStepRandomly(theGraph)

	}
	//allAnts
	theTravelPlan.Steps = counter
}
