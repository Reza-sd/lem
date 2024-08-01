package modelpk

import (
	// "fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
)

func ModelGeneratorA(myLem *Lem) (theModel Model) {

	NumberOfAnts := myLem.NumberOfAnts
	StartRoom := myLem.StartRoom
	EndRoom := myLem.EndRoom

	tunnelArr := myLem.TunnelArr

	var baseGraph graphpk.Graph
	baseGraph.InitGraph(tunnelArr, StartRoom, EndRoom)

	var baseAnts antpk.AntGroup
	baseAnts.AntsInit(NumberOfAnts, baseGraph.StartRoomName)

	theModel.BaseAnts = baseAnts
	theModel.BaseGraph = baseGraph
	return theModel
}
