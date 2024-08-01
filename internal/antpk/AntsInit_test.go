package antpk

import (
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
		assertIfAntsInitReturnError(t, err, expectError)

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
		assertIfAntsInitReturnError(t, err, expectError)

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
		assertIfAntsInitReturnError(t, err, expectError)

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
		assertIfAntsInitReturnError(t, err, expectError)

	})
}

// ================================================
func assertIfAntsInitReturnError(t testing.TB, err error, expectError bool) {
	t.Helper()

	if err == nil && expectError {
		t.Errorf("function returned error ")
	}

}

//================================================
