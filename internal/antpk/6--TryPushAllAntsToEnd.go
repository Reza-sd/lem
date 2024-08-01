package antpk

import (
	//"fmt"
	//"fmt"
	"fmt"
	"main/internal/graphpk"
)

// =====================TryPushAllAntsToEnd========================
func (allAnts *AntGroup) TryPushAllAntsToEnd(theGraph *graphpk.Graph, theTravelPlan *TravelPlan, maxTry int) {
	//counter := 0
	fmt.Println("UsedTunnelsInLastSequence", theGraph.UsedTunnelsInLastSequence)
	allAnts.NumberOfSequence = 0
	theGraph.UsedTunnelsInLastSequence = make(map[string]string)
	fmt.Println("0-UsedTunnelsInLastSequence", theGraph.UsedTunnelsInLastSequence)
	//--------------------------------
	for {
		//if all ants have arrived
		if theGraph.CurrentAntsInEndRoom == allAnts.NumberOfAnts {
			break
		}

		if allAnts.NumberOfSequence >= maxTry { // max try

			return
		}

		allAnts.MoveAllAntsOneStepRandomly(theGraph)
		//fmt.Println("allAnts.Step=",allAnts.Step)
		//counter++
		allAnts.NumberOfSequence++
		fmt.Println("counter=", allAnts.NumberOfSequence)
		//allAnts.Step=0

	}
	//---------------------------------
	// check if steps are less than previous steps
	if allAnts.NumberOfSequence >= theTravelPlan.FinalSequence && theTravelPlan.FinalSequence != 0 {
		return
	}

	theTravelPlan.FinalSequence = allAnts.NumberOfSequence
	theTravelPlan.TheBestPlan = *allAnts

}
