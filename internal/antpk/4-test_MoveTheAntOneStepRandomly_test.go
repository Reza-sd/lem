package antpk

import (
	//"reflect"
	"main/internal/graphpk"
	"testing"
	//"main/internal/simulationpk/modelpk"
)

// =======================================
func Test_MoveTheAntOneStepRandomly(t *testing.T) {
	//t.Skip()
	//------------
	t.Run(`1-return an error if baseAntGroup does not have any item`, func(t *testing.T) {
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
			SequenceNumber: 0,
			UsedTunnel:     map[int]map[string]string{},
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
	})
}
