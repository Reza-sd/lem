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
		//var secondAntGroup AntGroup
		var baseAntGroup AntGroup
		expectError := true
		//---Act---
		_, err := AntGroupCopyAtFirstRoom(baseAntGroup)
		//---Assert----
		assert_If_AntsCopy_ReturnError(t, err, expectError)
	})
	//------------
	t.Run(`2-return nil error if baseAntGroup have one item`, func(t *testing.T) {
		//---Arrange---
		//var secondAntGroup AntGroup
		//var baseAntGroup AntGroup
		baseAntGroup := Sample_AntGroup_1ant_initmode_room_0
		expectError := false //nil
		//---Act---
		_, err := AntGroupCopyAtFirstRoom(baseAntGroup)
		//---Assert----
		assert_If_AntsCopy_ReturnError(t, err, expectError)
	})
	//------------------
	t.Run(`3-return n same copy of baseAntGroup one item`, func(t *testing.T) {
		//---Arrange---
		baseAntGroup := Sample_AntGroup_1ant_initmode_room_0
		//var secondAntGroup AntGroup
		compareAntGroup := AntGroup{
			NumberOfAnts: 1,
			AntsDb: map[string]Ant{
				"L1": {Name: "L1",
					CurrentRoomName: "room_0",
					VisitedRoomsArr: []string{"room_0"},
					StepNumber:      0,
				},
			},
			currentSequenceNumber: 0,
			UsedTunnel:            make(map[int]map[string]string),
		}
		ifSame := true
		//---Act---
		secondAntGroup, _ := AntGroupCopyAtFirstRoom(baseAntGroup)
		//---Assert----
		assert_If_Two_AntGroup_SameCopy(t, compareAntGroup, secondAntGroup, ifSame)
	})
	//------------------
	t.Run(`4-return not same copy of baseAntGroup one item`, func(t *testing.T) {
		//---Arrange---
		baseAntGroup := Sample_AntGroup_1ant_initmode_room_0
		//var secondAntGroup AntGroup
		compareAntGroup := AntGroup{
			NumberOfAnts: 1,
			AntsDb: map[string]Ant{
				"L1": {Name: "L1",
					CurrentRoomName: "roooooom_0",
					VisitedRoomsArr: []string{"room_0"},
					StepNumber:      0,
				},
			},
			currentSequenceNumber: 0,
		}
		ifSame := false //not same
		//---Act---
		secondAntGroup, _ := AntGroupCopyAtFirstRoom(baseAntGroup)
		//---Assert----
		assert_If_Two_AntGroup_SameCopy(t, compareAntGroup, secondAntGroup, ifSame)
	})
	//------------------
	t.Run(`5-return n same copy of baseAntGroup two item`, func(t *testing.T) {
		//---Arrange---
		//var secondAntGroup AntGroup

		baseAntGroup := sample_AntGroup_2ants_initmode_room_0
		compareAntGroup := AntGroup{
			NumberOfAnts:          2,
			currentSequenceNumber: 0,
			AntsDb: map[string]Ant{
				"L1": {Name: "L1",
					CurrentRoomName: "room_0",
					VisitedRoomsArr: []string{"room_0"},
					StepNumber:      0,
				},
				"L2": {Name: "L2",
					CurrentRoomName: "room_0",
					VisitedRoomsArr: []string{"room_0"},
					StepNumber:      0,
				},
			},
			UsedTunnel: make(map[int]map[string]string),
		}
		ifSame := true
		//---Act---
		secondAntGroup, _ := AntGroupCopyAtFirstRoom(baseAntGroup)
		//---Assert----
		assert_If_Two_AntGroup_SameCopy(t, compareAntGroup, secondAntGroup, ifSame)
	})
	//--------------------------
}

// ================================================
func assert_If_Two_AntGroup_SameCopy(t testing.TB, antGroup1 AntGroup, antGroup2 AntGroup, ifSame bool) {
	t.Helper()

	if !reflect.DeepEqual(antGroup1, antGroup2) && ifSame {
		t.Errorf("not same \n antGroup1= %v \n antGroup2= %v ", antGroup1, antGroup2)
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
