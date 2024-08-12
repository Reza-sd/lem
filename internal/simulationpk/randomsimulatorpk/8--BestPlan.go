package randomsimulatorpk

import (
	//"fmt"
	//"fmt"
	//"fmt"
	//"main/internal/graphpk"
	"main/internal/antpk"
	"main/internal/simulationpk/modelpk"
)

func BestPlan(theModel *modelpk.Model) *TravelPlan {

	bestPlan := TravelPlan{
		FinalSequence: 10000,
		TheBestPlan:   antpk.AntGroup{},
	}

	for i := 1; i < 300; i++ {

		theModel.BaseAnts.TryPushAllAntsToEnd(&theModel.BaseGraph)
		if theModel.BaseAnts.CurrentSequenceNumber < bestPlan.FinalSequence && len(theModel.BaseAnts.NotArrivedAntsName) == 0 {
			bestPlan.FinalSequence = theModel.BaseAnts.CurrentSequenceNumber
			//(bestPlan.FinalSequence)
			bestPlan.TheBestPlan, _ = antpk.AntGroupCopyAtFirstRoom(theModel.BaseAnts)
		}
		theModel.ResetFactory()
	}
	//theModel.BaseAnts.TryPushAllAntsToEnd(&theModel.BaseGraph)

	return &bestPlan
}
