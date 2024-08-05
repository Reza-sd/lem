package antpk

import (
	//"fmt"
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
			CurrentSequenceNumber: 0,
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
		canIMove, _, err := CanImoveWhere("L100", myAntGroup, myGraph)
		//fmt.Println(">>>nextMove=", nextMove)
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
			CurrentSequenceNumber: 0,
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
		canIMove, _, err := CanImoveWhere("L1", myAntGroup, myGraph)
		//fmt.Println(">>>nextMove=", nextMove)
		assert_CanImoveWhere(t, err, expectError, canIMove, canIMoveExpected)
	})

	//------------
	t.Run("3-return nill error when ant and room exist", func(t *testing.T) {
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
		println("--->>> nextMove=", nextMove)
		assert_CanImoveWhere(t, err, expectError, canIMove, canIMoveExpected)
	})
	//------------
	t.Run("4-return err=nill and canmove=false=not-move when the ant already in end room", func(t *testing.T) {
		theAnt := Ant{
			Name:            "L1",
			CurrentRoomName: "room_1",
			VisitedRoomsArr: []string{"room_1"},
			StepNumber:      0,
		}
		myAntGroup := AntGroup{
			NumberOfAnts: 1,
			AntsDb: map[string]Ant{
				"L1": theAnt,
			},
			CurrentSequenceNumber: 0,
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
		canIMoveExpected := false
		canIMove, _, err := CanImoveWhere("L1", myAntGroup, myGraph)
		//fmt.Println(">>>nextMove=", nextMove)
		assert_CanImoveWhere(t, err, expectError, canIMove, canIMoveExpected)
	})
	//------------
	t.Run("5-return err=nill and canmove=false=not-move when the next room dose not have empty seat", func(t *testing.T) {
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
			UsedTunnel:            map[int]map[string]string{},
		}
		myGraph := graphpk.Graph{
			StartRoomName:        "room_0",
			EndRoomName:          "room_1",
			CurrentAntsInEndRoom: 0,
			NumberOfAllRoom:      2,
			Rooms: map[string]graphpk.Room{
				"room_0": graphpk.Sample_room_0_startRoom,
				"room_1": graphpk.Room{
					Name:        "room_1",
					MaxSeats:    1000,
					EmptySeats:  0, //<----------0
					Connections: []string{"room_0"},
				},
			},
		}
		expectError := false //nil
		canIMoveExpected := false
		canIMove, _, err := CanImoveWhere("L1", myAntGroup, myGraph)
		//fmt.Println(">>>nextMove=", nextMove)
		assert_CanImoveWhere(t, err, expectError, canIMove, canIMoveExpected)
	})
	//------------
	t.Run("6-return err=nill and canmove=false=not-move when offered tunnel used in this sequence", func(t *testing.T) {
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
			CurrentSequenceNumber: 1,
			UsedTunnel: map[int]map[string]string{
				1: {
					"room_0": "room_1",
				},
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
		expectError := false //nil
		canIMoveExpected := false
		canIMove, _, err := CanImoveWhere("L1", myAntGroup, myGraph)
		//fmt.Println(">>>nextMove=", nextMove)
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

	} else {
		if err != nil {
			t.Errorf(" return= %v expect= %v", err, expectError)
		}

	}

	if canIMove != canIMoveExpected {
		t.Errorf(">>(canIMove= %v, canIMoveExpected= %v)<<", canIMove, canIMoveExpected)

	}
}

//=================================
