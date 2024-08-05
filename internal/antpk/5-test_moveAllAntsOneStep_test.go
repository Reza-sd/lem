package antpk

import (
	"main/internal/graphpk"
	"testing"
)

func Test_MoveAllAntsOneStepRandomly(t *testing.T) {
	//-----------------------------------
	t.Run("1-", func(t *testing.T) {
		t.Skip()
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
			CurrentSequenceNumber: 0,
			UsedTunnel:            make(map[int]map[string]string),

			NotArrivedAntsName: []string{"L1"},
		}

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
		//myGraph.Print()
		myAntGroup.MoveAllAntsOneStepRandomly(&myGraph)
		myAntGroup.Print()
		//myGraph.Print()
		myAntGroup.MoveAllAntsOneStepRandomly(&myGraph)
		myAntGroup.Print()
		myAntGroup.MoveAllAntsOneStepRandomly(&myGraph)
		myAntGroup.Print()
		//myGraph.Print()
	})
	//-----------------------------------
	t.Run("2-", func(t *testing.T) {
		//t.Skip()
		ant1 := Ant{
			Name:            "L1",
			CurrentRoomName: "room_0",
			VisitedRoomsArr: []string{"room_0"},
			StepNumber:      0,
		}
		ant2 := Ant{
			Name:            "L2",
			CurrentRoomName: "room_0",
			VisitedRoomsArr: []string{"room_0"},
			StepNumber:      0,
		}
		myAntGroup := AntGroup{
			NumberOfAnts: 2,
			AntsDb: map[string]Ant{
				"L1": ant1,
				"L2": ant2,
			},
			CurrentSequenceNumber: 0,
			UsedTunnel:            make(map[int]map[string]string),

			NotArrivedAntsName: []string{"L1","L2"},
		}

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
		//myGraph.Print()
		myAntGroup.MoveAllAntsOneStepRandomly(&myGraph)
		myAntGroup.Print()
		myAntGroup.MoveAllAntsOneStepRandomly(&myGraph)
		myAntGroup.Print()
		//myGraph.Print()
	})
}
