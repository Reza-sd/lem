package modelpk

import (
	// "fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
)

func ModelGeneratorA(myLem Lem) (theModel Model) {

	//myLem.baseModel.baseAnts=

	NumberOfAnts := myLem.NumberOfAnts
	StartRoom := myLem.StartRoom
	EndRoom := myLem.EndRoom

	tunnelArr := myLem.TunnelArr

	var baseGraph graphpk.Graph
	baseGraph.InitGraph(tunnelArr, StartRoom, EndRoom)

	var baseAnts antpk.Ants
	baseAnts.AntsInit(NumberOfAnts, &baseGraph)

	theModel.BaseAnts = baseAnts
	theModel.BaseGraph = baseGraph
	return theModel
}
