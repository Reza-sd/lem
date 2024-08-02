package antpk

import (
	"reflect"
	"testing"
)

func Test_AntsCopy(t *testing.T) {
	//t.Skip()
	//------------
	t.Run(`1-return an error if baseAntGroup does not have any item`, func(t *testing.T) {
		//---Arrange---
		var secondAntGroup AntGroup
		var baseAntGroup AntGroup
		expectError := true
		//---Act---
		err := AntGroupCopyAtFirstRoom(baseAntGroup, &secondAntGroup)
		//---Assert----
		assert_If_AntsCopy_ReturnError(t, err, expectError)
	})
	//------------
	t.Run(`2-return nil error if baseAntGroup have one item`, func(t *testing.T) {
		//---Arrange---
		var secondAntGroup AntGroup
		//var baseAntGroup AntGroup
		baseAntGroup := Sample_AntGroup_1ant_initmode_room_0
		expectError := false //nil
		//---Act---
		err := AntGroupCopyAtFirstRoom(baseAntGroup, &secondAntGroup)
		//---Assert----
		assert_If_AntsCopy_ReturnError(t, err, expectError)
	})
	//------------------
	t.Run(`3-return n same copy of baseAntGroup one item`, func(t *testing.T) {
		//---Arrange---
		baseAntGroup := Sample_AntGroup_1ant_initmode_room_0
		var secondAntGroup AntGroup
		compareAntGroup := AntGroup{
			NumberOfAnts: 1,
			AntsMap: map[int]Ant{
				1: {Name: "L1",
					CurrentRoomName: "room_0",
					VisitedRoomsArr: []string{"room_0"},
					StepNumber:      0,
				},
			},
			SequenceNumber: 0,
			UsedTunnel: TravelHistory{
				UsedTunnelsMap: make(map[int]map[string]string),
			},
		}
		ifSame := true
		//---Act---
		AntGroupCopyAtFirstRoom(baseAntGroup, &secondAntGroup)
		//---Assert----
		assert_If_Two_AntGroup_SameCopy(t, compareAntGroup, secondAntGroup, ifSame)
	})
	//------------------
	t.Run(`4-return not same copy of baseAntGroup one item`, func(t *testing.T) {
		//---Arrange---
		baseAntGroup := Sample_AntGroup_1ant_initmode_room_0
		var secondAntGroup AntGroup
		compareAntGroup := AntGroup{
			NumberOfAnts: 1,
			AntsMap: map[int]Ant{
				1: {Name: "L1",
					CurrentRoomName: "roooooom_0",
					VisitedRoomsArr: []string{"room_0"},
					StepNumber:      0,
				},
			},
			SequenceNumber: 0,
		}
		ifSame := false //not same
		//---Act---
		AntGroupCopyAtFirstRoom(baseAntGroup, &secondAntGroup)
		//---Assert----
		assert_If_Two_AntGroup_SameCopy(t, compareAntGroup, secondAntGroup, ifSame)
	})
	//------------------
	t.Run(`5-return n same copy of baseAntGroup two item`, func(t *testing.T) {
		//---Arrange---
		var secondAntGroup AntGroup

		baseAntGroup := sample_AntGroup_2ants_initmode_room_0
		compareAntGroup := AntGroup{
			NumberOfAnts:   2,
			SequenceNumber: 0,
			AntsMap: map[int]Ant{
				1: {Name: "L1",
					CurrentRoomName: "room_0",
					VisitedRoomsArr: []string{"room_0"},
					StepNumber:      0,
				},
				2: {Name: "L2",
					CurrentRoomName: "room_0",
					VisitedRoomsArr: []string{"room_0"},
					StepNumber:      0,
				},
			},
			UsedTunnel: TravelHistory{
				UsedTunnelsMap: make(map[int]map[string]string),
			},
		}
		ifSame := true
		//---Act---
		AntGroupCopyAtFirstRoom(baseAntGroup, &secondAntGroup)
		//---Assert----
		assert_If_Two_AntGroup_SameCopy(t, compareAntGroup, secondAntGroup, ifSame)
	})
	//--------------------------
}

// ================================================
func assert_If_Two_AntGroup_SameCopy(t testing.TB, antGroup1 AntGroup, antGroup2 AntGroup, ifSame bool) {
	t.Helper()

	if !reflect.DeepEqual(antGroup1, antGroup2) && ifSame {
		t.Errorf("function returned error ")
	}

}

// ================================================
func assert_If_AntsCopy_ReturnError(t testing.TB, err error, expectError bool) {
	t.Helper()

	if err == nil && expectError {
		t.Errorf("function returned error ")
	}

}

// ================================================
