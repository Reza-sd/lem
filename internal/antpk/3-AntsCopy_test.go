package antpk

import (
	"reflect"
	"testing"
)

func Test_AntsCopy(t *testing.T) {
	//------------
	t.Run(`0-return an error if baseAntGroup does not have any item`, func(t *testing.T) {
		//---Arrange---
		var secondAntGroup AntGroup
		var baseAntGroup AntGroup
		expectError := true
		//---Act---
		err := AntGroupCopy(baseAntGroup, &secondAntGroup)
		//---Assert----
		assert_If_AntsCopy_ReturnError(t, err, expectError)
	})
	//------------
	t.Run(`1-return nil error if baseAntGroup have one item`, func(t *testing.T) {
		//---Arrange---
		var secondAntGroup AntGroup
		//var baseAntGroup AntGroup
		baseAntGroup := Sample_AntGroup_1ant_initmode_room_0
		expectError := false //nil
		//---Act---
		err := AntGroupCopy(baseAntGroup, &secondAntGroup)
		//---Assert----
		assert_If_AntsCopy_ReturnError(t, err, expectError)
	})
	//------------------
	t.Run(`2-return n same copy of baseAntGroup`, func(t *testing.T) {
		//---Arrange---
		//var secondAntGroup AntGroup
		//var baseAntGroup AntGroup
		baseAntGroup := Sample_AntGroup_1ant_initmode_room_0
		secondAntGroup := AntGroup{
			NumberOfAnts: 1,
			AntsMap: map[int]Ant{
				1: {Name: "L1",
					CurrentRoomName: "room_0",
					VisitedRoomsArr: []string{"room_0"},
					StepNumber:      0,
				},
			},
			NumberOfSequence: 0,
		}
		//expectError := true
		//---Act---
		AntGroupCopy(baseAntGroup, &secondAntGroup)
		//---Assert----
		assert_If_Two_AntGroup_SameCopy(t, baseAntGroup, secondAntGroup)
	})
}

// ================================================
func assert_If_Two_AntGroup_SameCopy(t testing.TB, antGroup1 AntGroup, antGroup2 AntGroup) {
	t.Helper()

	if !reflect.DeepEqual(antGroup1, antGroup2) {
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
