package antpk

import (
	"fmt"
	"main/internal/graphpk"
	"testing"
)

// =======================================
func Test_CanImoveWhere(t *testing.T) {
	//t.Skip()
	//------------
	t.Run("1-return error when ant does not exist", func(t *testing.T) {
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

		expectError := true
		canIMoveExpected := false
		canIMove, nextMove, err := CanImoveWhere("L1d", myAntGroup, myGraph)
		fmt.Println(">>>nextMove=", nextMove)
		assert_CanImoveWhere(t, err, expectError, canIMove, canIMoveExpected)
	})
	//------------
	t.Run("2-return error when roon does not exist", func(t *testing.T) {
		theAnt := Ant{
			Name:            "L1",
			CurrentRoomName: "room_3",
			VisitedRoomsArr: []string{"room_3"},
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

		expectError := true
		canIMoveExpected := false
		canIMove, nextMove, err := CanImoveWhere("L1", myAntGroup, myGraph)
		fmt.Println(">>>nextMove=", nextMove)
		assert_CanImoveWhere(t, err, expectError, canIMove, canIMoveExpected)
	})
	//------------
	t.Run("3-return error when roon does not exist", func(t *testing.T) {
		theAnt := Ant{
			Name:            "L1",
			CurrentRoomName: "room_3",
			VisitedRoomsArr: []string{"room_3"},
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

		expectError := true
		canIMoveExpected := false
		canIMove, nextMove, err := CanImoveWhere("L1", myAntGroup, myGraph)
		fmt.Println(">>>nextMove=", nextMove)
		assert_CanImoveWhere(t, err, expectError, canIMove, canIMoveExpected)
	})
	//------------
	t.Run("4-return nill error when ant and room exist", func(t *testing.T) {
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
		expectError := false //nil
		canIMoveExpected := true
		canIMove, nextMove, err := CanImoveWhere("L1", myAntGroup, myGraph)
		fmt.Println(">>>nextMove=", nextMove)
		assert_CanImoveWhere(t, err, expectError, canIMove, canIMoveExpected)
	})
}

// =================================
func assert_CanImoveWhere(t testing.TB, err error, expectError bool, canIMove bool, canIMoveExpected bool) {
	t.Helper()

	if expectError {
		if err == nil {
			t.Errorf(" return= %v expect= %v", err, expectError)

		}
		//  else {
		// 	if nextMove != "" {
		// 		t.Errorf("function return nextMove= %v expect Empty", nextMove)
		// 	}
		// }
	} else {
		if err != nil {
			t.Errorf(" return= %v expect= %v", err, expectError)
		}
		// else {
		// 	if nextMove == expectedNextMove {
		// 		t.Errorf("function return nextMove= %v expect %v", expectedNextMove)
		// 	}
		// }

	}

	if canIMove != canIMoveExpected {
		t.Errorf(">>(canIMove= %v, canIMoveExpected= %v)<<", canIMove, canIMoveExpected)

	}
}

//=================================
