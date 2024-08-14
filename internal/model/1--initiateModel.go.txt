package modelpk

import (
	"fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
)

func ModelGeneratorA(myLem *Lem) (Model, error) {
	var theModel Model
	NumberOfAnts := myLem.NumberOfAnts
	StartRoom := myLem.StartRoom
	EndRoom := myLem.EndRoom

	tunnelArr := myLem.TunnelArr

	var baseGraph graphpk.Graph
	baseGraph.InitGraph(tunnelArr, StartRoom, EndRoom)

	var baseAnts antpk.AntGroup
	errBaseAntsAntsInit := baseAnts.AntGroupInit(NumberOfAnts, baseGraph.StartRoomName)

	if errBaseAntsAntsInit != nil {
		return theModel, fmt.Errorf("number of Ants is not in [1-200]")
	}
	theModel.BaseAnts = baseAnts
	theModel.BaseGraph = baseGraph
	return theModel, nil
}
