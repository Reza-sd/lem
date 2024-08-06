package antpk

import (
	//"fmt"
	//"fmt"
	//"fmt"
	"main/internal/graphpk"
)

// =====================TryPushAllAntsToEnd========================
func (allAnts *AntGroup) TryPushAllAntsToEnd(theGraph *graphpk.Graph) {
	//counter:=0
	for i := 1; i < 20; i++ {

		if len(allAnts.NotArrivedAntsName) == 0 {
			break
		}
		allAnts.TryMoveAllAntsOneStepRandomly(theGraph)

	}

}

/*
func (allAnts *AntGroup) TryPushAllAntsToEnd(theGraph *graphpk.Graph, theTravelPlan *TravelPlan, maxTry int, travelHistory *TravelHistory) {
	//counter := 0
	//fmt.Println("UsedTunnelsInLastSequence", theGraph.UsedTunnelsInLastSequence)
	allAnts.SequenceNumber = 0
	//theGraph.UsedTunnelsInLastSequence = make(map[string]string)
	//fmt.Println("0-UsedTunnelsInLastSequence", theGraph.UsedTunnelsInLastSequence)
	//--------------------------------
	for {
		//if all ants have arrived
		if theGraph.CurrentAntsInEndRoom == allAnts.NumberOfAnts {
			break
		}

		if allAnts.SequenceNumber >= maxTry { // max try

			return
		}

		allAnts.MoveAllAntsOneStepRandomly(theGraph, travelHistory)
		//fmt.Println("allAnts.Step=",allAnts.Step)
		//counter++
		allAnts.SequenceNumber++
		fmt.Println("counter=", allAnts.SequenceNumber)
		//allAnts.Step=0

	}
	//---------------------------------
	// check if steps are less than previous steps
	if allAnts.SequenceNumber >= theTravelPlan.FinalSequence && theTravelPlan.FinalSequence != 0 {
		return
	}

	theTravelPlan.FinalSequence = allAnts.SequenceNumber
	theTravelPlan.TheBestPlan = *allAnts

}
*/
