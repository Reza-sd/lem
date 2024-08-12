package antpk

import (
	"reflect"
	"testing"
)

// ========================================

// =========================================
func Test_AntsInit(t *testing.T) {
	//t.Skip()
	//------------
	t.Run(`1-return an error when numberOfAnts=0`, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := mta(0)

		expectError := true
		//---Act---
		err := myAntGroup.AntGroupInit(numberOfAnts)
		//---Assert----
		assert_If_AntsInit_ReturnError(t, err, expectError)

	})

	// 	//------------
	t.Run(`2-return an error when numberOfAnts > MaxHandleableAntsNumber`, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := MaxHandleableAntsNumber + 1

		expectError := true
		//---Act---
		err := myAntGroup.AntGroupInit(numberOfAnts)
		//---Assert----
		assert_If_AntsInit_ReturnError(t, err, expectError)

	})

	// 	//------------
	t.Run(`3-return nil error when startRoomName=0 and numberOfAnts>0 `, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := mta(1)

		expectError := false //no error = nil
		//---Act---
		err := myAntGroup.AntGroupInit(numberOfAnts)
		//---Assert----
		assert_If_AntsInit_ReturnError(t, err, expectError)

	})
	// 	//------------
	t.Run(`4-check if the values of myAntGroup correct (1 ant) `, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := mta(1)

		expectedAntGroup := Sample_AntGroup_1ant_initmode_room_0
		//---Act---
		myAntGroup.AntGroupInit(numberOfAnts)
		//---Assert----
		assert_If_Two_AntGroup_Same(t, myAntGroup, expectedAntGroup)

	})
	// 	//------------
	t.Run(`5-check if the values of myAntGroup correct (2 ant)`, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		numberOfAnts := mta(2)

		expectedAntGroup := Sample_AntGroup_2ant_initmode_room_0
		//---Act---
		myAntGroup.AntGroupInit(numberOfAnts)
		//---Assert----
		assert_If_Two_AntGroup_Same(t, myAntGroup, expectedAntGroup)
		//myAntGroup.Print()

	})
}

// ================================================
func assert_If_Two_AntGroup_Same(t testing.TB, antGroup1 AntGroup, expectedAntGroup AntGroup) {
	t.Helper()

	if !reflect.DeepEqual(antGroup1, expectedAntGroup) {
		t.Errorf("\n>>>not same: \n antGroup1=%v<<\n expectedAntGroup=%v<<", antGroup1, expectedAntGroup)
	}

}

// ================================================
func assert_If_AntsInit_ReturnError(t testing.TB, err error, expectError bool) {
	t.Helper()

	if expectError {
		if err == nil {
			t.Errorf("function return %v expect %v", err, expectError)

		}
	} else {
		if err != nil {
			t.Errorf("function return %v expect %v", err, expectError)
		}
	}

}

//================================================
