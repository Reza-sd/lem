package antpk

import (
	
	"fmt"
	"main/internal/graphpk"
	"testing"

)

// =======================================
func Test_WhatsMyNextMove(t *testing.T) {
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
		nextMove, err := WhatsMyNextMove("L1d", myAntGroup, myGraph)
		fmt.Println(">>>nextMove=", nextMove)
		assert_WhatsMyNextMove(t, err, expectError, nextMove)
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
		nextMove, err := WhatsMyNextMove("L1", myAntGroup, myGraph)
		fmt.Println(">>>nextMove=", nextMove)
		assert_WhatsMyNextMove(t, err, expectError, nextMove)
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
		nextMove, err := WhatsMyNextMove("L1", myAntGroup, myGraph)
		fmt.Println(">>>nextMove=", nextMove)
		assert_WhatsMyNextMove(t, err, expectError, nextMove)
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
		nextMove, err := WhatsMyNextMove("L1", myAntGroup, myGraph)
		fmt.Println(">>>nextMove=", nextMove)
		assert_WhatsMyNextMove(t, err, expectError, nextMove)
	})
}

// =================================
func assert_WhatsMyNextMove(t testing.TB, err error, expectError bool, nextMove string) {
	t.Helper()

	if expectError {
		if err == nil {
			t.Errorf("function return %v expect %v", err, expectError)

		} else {
			if nextMove != "" {
				t.Errorf("function return nextMove= %v expect Empty", nextMove)
			}
		}
	} else {
		if err != nil {
			t.Errorf("function return %v expect %v", err, expectError)
		} else {
			if nextMove == "" {
				t.Errorf("function return nextMove= %v expect Empty", nextMove)
			}
		}

	}

}

//=================================
