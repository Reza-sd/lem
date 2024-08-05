package antpk

import (
	//"reflect"
	"fmt"
	"main/internal/graphpk"
	"testing"
	//"main/internal/simulationpk/modelpk"
)

// =======================================
func Test_WhatsMyNextMove(t *testing.T) {
	//t.Skip()
	//------------
	t.Run(`1-return `, func(t *testing.T) {
		theAnt := Ant{
			Name:            "L1",
			CurrentRoomName: "room_0",
			VisitedRoomsArr: []string{"room_0"},
			StepNumber:      0,
		}
		myAntGroup := AntGroup{
			NumberOfAnts: 1,
			AntsDb: map[string]Ant{
				"L1": theAnt,
			},
			currentSequenceNumber: 0,
			UsedTunnel:            map[int]map[string]string{},
		}
		//travelHistory :=myAntGroup.UsedTunnel
		myGraph := graphpk.Graph{
			StartRoomName:        "room_0",
			EndRoomName:          "room_1",
			CurrentAntsInEndRoom: 0,
			NumberOfAllRoom:      2,
			Rooms: map[string]graphpk.Room{
				"room_0": graphpk.Sample_room_0_startRoom,
				"room_1": graphpk.Sample_room_1_EndRoom,
			},
		}

		myAntGroup.Print()
		myGraph.Print()
		expectError := true
		nextMove, err := WhatsMyNextMove("L1d", myAntGroup, myGraph)
		fmt.Println(">>>nextMove=", nextMove)
		assert_WhatsMyNextMove(t, err, expectError)
	})
}

// =================================
func assert_WhatsMyNextMove(t testing.TB, err error, expectError bool) {
	t.Helper()

	if expectError {
		if err == nil {
			t.Errorf("function return %v expect %v", err, expectError)

		}
	} else {
		if err != nil {
			t.Errorf("function return %v expect %v", err, expectError)
		}
	}

}

//=================================
