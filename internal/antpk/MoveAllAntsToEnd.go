package antpk

import (
	//"fmt"
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *Ants) PushAllAntsToEnd(theGraph *graphpk.Graph, theTravelPlan *TravelPlan) {
	counter := 0
	//success := true
	for {

		if theGraph.CurrentAntsInEndRoom == allAnts.AntsNumber {
			break
		}

		if counter >= 10 { // max try
			//success = false
			//break
			return
		}

		allAnts.MoveAllAntsOneStepRandomly(theGraph)
		counter++
	}
	//allAnts
	//if !success { return}
	if counter >= theTravelPlan.Steps || counter == 0 {
		return
	}
	// fmt.Println("counter=",counter)
	// fmt.Println("theTravelPlan.Steps=",theTravelPlan.Steps)

	if counter != 0 {
		theTravelPlan.Steps = counter
		theTravelPlan.SuccessfulPlan = *allAnts
	}

}
