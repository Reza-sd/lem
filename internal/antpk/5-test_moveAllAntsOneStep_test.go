package antpk

import (
	"main/internal/graphpk"
	"reflect"
	"testing"
)

func Test_MoveAllAntsOneStepRandomly(t *testing.T) {

	//-----------------------------------
	t.Run("1-move 1 ant", func(t *testing.T) {
		//t.Skip()
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
			UsedTunnel:            make(map[string][]string),

			NotArrivedAntsName: map[string]struct{}{"L1": {}},
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
		myAntGroup.TryMoveAllAntsOneStepRandomly(&myGraph)
		myAntGroup.Print()
		expSequenceNumbebr := 1
		expUsedTunnel := map[int]map[string]string{
			1: {
				"room_0": "room_1",
			},
		}
		//expNotArrivedAntsName := map[string]struct{}{}
		assert_MoveAllAntsOneStepRandomly(t, &myAntGroup, expSequenceNumbebr, expUsedTunnel)
		//myGraph.Print()
		// myAntGroup.MoveAllAntsOneStepRandomly(&myGraph)
		// myAntGroup.Print()
		// myAntGroup.MoveAllAntsOneStepRandomly(&myGraph)
		// myAntGroup.Print()
		//myGraph.Print()
	})
	//-----------------------------------
	t.Run("2-move 2 ant", func(t *testing.T) {
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
			UsedTunnel:            make(map[string][]string),

			NotArrivedAntsName: map[string]struct{}{
				"L1": {},
				"L2": {},
			},
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
		myAntGroup.TryMoveAllAntsOneStepRandomly(&myGraph)
		expSequenceNumbebr := 1
		expUsedTunnel := map[int]map[string]string{
			1: {
				"room_0": "room_1",
			},
		}
		// expNotArrivedAntsName := map[string]struct{}{
		// 	"L1": struct{}{},
		// }
		assert_MoveAllAntsOneStepRandomly(t, &myAntGroup, expSequenceNumbebr, expUsedTunnel)

		myAntGroup.Print()
		// myAntGroup.TryMoveAllAntsOneStepRandomly(&myGraph)

		// expSequenceNumbebr = 2
		// expUsedTunnel = map[int]map[string]string{
		// 	1: {
		// 		"room_0": "room_1",
		// 	},
		// 	2: {
		// 		"room_0": "room_1",
		// 	},
		// }
		// expNotArrivedAntsName = map[string]bool{"L1":true}
		// assert_MoveAllAntsOneStepRandomly(t, &myAntGroup, expSequenceNumbebr, expUsedTunnel, expNotArrivedAntsName)
		// myAntGroup.Print()
		// myAntGroup.MoveAllAntsOneStepRandomly(&myGraph)
		// myAntGroup.Print()
		//myGraph.Print()
	})
	//-----------------------------------
	t.Run("3-move 3 ant", func(t *testing.T) {
		//.Skip()
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
		ant3 := Ant{
			Name:            "L3",
			CurrentRoomName: "room_0",
			VisitedRoomsArr: []string{"room_0"},
			StepNumber:      0,
		}

		myAntGroup := AntGroup{
			NumberOfAnts: 3,
			AntsDb: map[string]Ant{
				"L1": ant1,
				"L2": ant2,
				"L3": ant3,
			},
			CurrentSequenceNumber: 0,
			UsedTunnel:            make(map[string][]string),

			NotArrivedAntsName: map[string]struct{}{
				"L1": {},
				"L2": {},
				"L3": {},
			},
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
		myAntGroup.TryMoveAllAntsOneStepRandomly(&myGraph)
		expSequenceNumbebr := 1
		expUsedTunnel := map[int]map[string]string{
			1: {
				"room_0": "room_1",
			},
		}
		// expNotArrivedAntsName := map[string]struct{}{
		// 	"L1": struct{}{},
		// }
		myAntGroup.Print()

		assert_MoveAllAntsOneStepRandomly(t, &myAntGroup, expSequenceNumbebr, expUsedTunnel)

		// myAntGroup.Print()
		// myAntGroup.MoveAllAntsOneStepRandomly(&myGraph)

		// expSequenceNumbebr = 2
		// expUsedTunnel = map[int]map[string]string{
		// 	1: {
		// 		"room_0": "room_1",
		// 	},
		// 	2: {
		// 		"room_0": "room_1",
		// 	},
		// }
		// expNotArrivedAntsName = []string{}
		// assert_MoveAllAntsOneStepRandomly(t, &myAntGroup, expSequenceNumbebr, expUsedTunnel, expNotArrivedAntsName)
		// myAntGroup.Print()
		// myAntGroup.MoveAllAntsOneStepRandomly(&myGraph)
		// myAntGroup.Print()
		//myGraph.Print()
	})
}

// ======================================
func assert_MoveAllAntsOneStepRandomly(t testing.TB, myAntGroup *AntGroup, expSequenceNumbebr int, expUsedTunnel map[int]map[string]string) {
	if expSequenceNumbebr != myAntGroup.CurrentSequenceNumber {
		t.Errorf("exp %v receive %v", expSequenceNumbebr, myAntGroup.CurrentSequenceNumber)
	}

	if !reflect.DeepEqual(expUsedTunnel, myAntGroup.UsedTunnel) {
		t.Errorf("exp %v receive %v", expUsedTunnel, myAntGroup.UsedTunnel)
	}

	// if !reflect.DeepEqual(expNotArrivedAntsName, myAntGroup.NotArrivedAntsName) {
	// 	t.Errorf("exp %v receive %v", expNotArrivedAntsName, myAntGroup.NotArrivedAntsName)
	// }

}

//============================================
