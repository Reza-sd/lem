package modelpk

import (
	antpk "main/internal/model/antgroup"
	//graphpk "main/internal/model/graphpk"
)

func (myModel *Model) Init(myLem *Lem) error {

	myModel.AntGroup.Init(antpk.Mtag(myLem.NumberOfAnts))
	myModel.Graph.Init(myLem.TunnelMap, antpk.Mtag(myLem.EndRoom))
	// var theModel Model
	// NumberOfAnts := myLem.NumberOfAnts
	// StartRoom := myLem.StartRoom
	// EndRoom := myLem.EndRoom

	// tunnelArr := myLem.TunnelArr

	// var baseGraph graphpk.Graph
	// baseGraph.InitGraph(tunnelArr, StartRoom, EndRoom)

	// var baseAnts antpk.AntGroup
	// errBaseAntsAntsInit := baseAnts.AntGroupInit(NumberOfAnts, baseGraph.StartRoomName)

	// if errBaseAntsAntsInit != nil {
	// 	return theModel, fmt.Errorf("number of Ants is not in [1-200]")
	// }
	// theModel.BaseAnts = baseAnts
	// theModel.BaseGraph = baseGraph
	return nil
}
