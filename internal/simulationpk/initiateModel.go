package simulationpk

import (
	// "fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
)

func (myLem *Lem) ModelInit() (theModel Model){

	//myLem.baseModel.baseAnts=

	NumberOfAnts := myLem.NumberOfAnts
	StartRoom := myLem.StartRoom
	EndRoom := myLem.EndRoom

	tunnelArr := myLem.TunnelArr

	var baseGraph graphpk.Graph
	baseGraph.InitGraph(tunnelArr, StartRoom, EndRoom)

	var baseAnts antpk.Ants
	baseAnts.AntsInit(NumberOfAnts, &baseGraph)

	theModel.baseAnts = baseAnts
	theModel.baseGraph = baseGraph
	return theModel
}
