package modelpk

import (
	//"fmt"
	antpk "main/internal/model/antpk"
	graphpk "main/internal/model/graphpk"
)

var (
	myGraph    = graphpk.Sample_graph_1
	myAntGroup = antpk.Sample_AntGroup_2ant_initmode_room_0

	myModel = Model{
		Graph:    myGraph,
		AntGroup: myAntGroup,
	}
)

//-------------------------
