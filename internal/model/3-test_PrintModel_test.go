package modelpk

import (
	"testing"
)

func Test_Print(t *testing.T) {

	t.Run(`1-Print`, func(t *testing.T) {

		// myGraph = graphpk.Sample_graph_1
		// myAntGroup = antpk.Sample_AntGroup_2ant_initmode_room_0

		// myModel = Model{
		// 	Graph : myGraph,
		// 	AntGroup :myAntGroup,
		// }
		myModel.Print()

	})
}
