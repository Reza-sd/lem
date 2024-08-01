package antpk

import (
	"reflect"
	"testing"
)

// ========================================
var ()

// =========================================
func Test_AntsInit(t *testing.T) {
	//------------
	t.Run(`0-return an error when numberOfAnts=0`, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := 0
		startRoomName := "L1"
		expectError := true
		//---Act---
		err := myAntGroup.AntsInit(numberOfAnts, startRoomName)
		//---Assert----
		assert_If_AntsInit_ReturnError(t, err, expectError)

	})
	//------------
	t.Run(`001-return an error when numberOfAnts< 0`, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := -2
		startRoomName := "L1"
		expectError := true
		//---Act---
		err := myAntGroup.AntsInit(numberOfAnts, startRoomName)
		//---Assert----
		assert_If_AntsInit_ReturnError(t, err, expectError)

	})
	//------------
	t.Run(`0-1-return an error when numberOfAnts > MaxHandleableAntsNumber`, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := MaxHandleableAntsNumber + 1
		startRoomName := "L1"
		expectError := true
		//---Act---
		err := myAntGroup.AntsInit(numberOfAnts, startRoomName)
		//---Assert----
		assert_If_AntsInit_ReturnError(t, err, expectError)

	})
	//------------
	t.Run(`1-return an error when startRoomName="" `, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := 1
		startRoomName := ""
		expectError := true
		//---Act---
		err := myAntGroup.AntsInit(numberOfAnts, startRoomName)
		//---Assert----
		assert_If_AntsInit_ReturnError(t, err, expectError)

	})
	//------------
	t.Run(`2-return nil error when startRoomName="L1" and numberOfAnts>0 `, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := 1
		startRoomName := "L1"
		expectError := false //no error = nil
		//---Act---
		err := myAntGroup.AntsInit(numberOfAnts, startRoomName)
		//---Assert----
		assert_If_AntsInit_ReturnError(t, err, expectError)

	})
	//------------
	t.Run(`3-check if the values of myAntGroup correct (1 ant) `, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := 1
		startRoomName := "room_4"
		//expectError := false //no error = nil
		expectedmyAntGroup := AntGroup{
			NumberOfAnts: 1,
			AntsMap: map[int]Ant{
				1: {Name: "L1",
					CurrentRoomName: "room_4",
					VisitedRoomsArr: []string{"room_4"},
					StepNumber:      0,
				},
			},
			NumberOfSequence: 0,
		}
		//---Act---
		myAntGroup.AntsInit(numberOfAnts, startRoomName)
		//---Assert----
		assert_If_Two_AntGroup_Same(t, myAntGroup, expectedmyAntGroup)

	})
	//------------
	t.Run(`4-check if the values of myAntGroup correct (2 ant)`, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := 2
		startRoomName := "room_4"
		//expectError := false //no error = nil
		expectedmyAntGroup := AntGroup{
			NumberOfAnts: 2,
			AntsMap: map[int]Ant{
				1: {Name: "L1",
					CurrentRoomName: "room_4",
					VisitedRoomsArr: []string{"room_4"},
					StepNumber:      0,
				},
				2: {Name: "L2",
					CurrentRoomName: "room_4",
					VisitedRoomsArr: []string{"room_4"},
					StepNumber:      0,
				},
			},
			NumberOfSequence: 0,
		}
		//---Act---
		myAntGroup.AntsInit(numberOfAnts, startRoomName)
		//---Assert----
		assert_If_Two_AntGroup_Same(t, myAntGroup, expectedmyAntGroup)

	})
}

// ================================================
func assert_If_Two_AntGroup_Same(t testing.TB, antGroup1 AntGroup, antGroup2 AntGroup) {
	t.Helper()

	if !reflect.DeepEqual(antGroup1, antGroup2) {
		t.Errorf("function returned error ")
	}

}

// ================================================
func assert_If_AntsInit_ReturnError(t testing.TB, err error, expectError bool) {
	t.Helper()

	if err == nil && expectError {
		t.Errorf("function returned error ")
	}

}

//================================================
