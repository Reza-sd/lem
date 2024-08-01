package antpk

import (
	"reflect"
	"testing"
)

// ===============================================
/*

 */

// ===============================================
var(
	//-------Ant type example-----------
	ant1 = Ant{
		Name:            "reza",
		CurrentRoomName: "room_1",
		VisitedRoomsArr: []string{"room_0", "room_1"},
		StepNumber:      2,
	}
	//--------AntGroup type example---------
	AntGroup1 = AntGroup{
		NumberOfAnts: 1,
		AntsMap: map[int]Ant{
			1: {Name: "Ant1",
				CurrentRoomName: "room_1",
				VisitedRoomsArr: []string{"room_0", "room_1"},
				StepNumber:      2,
			},
		},
		NumberOfSequence: 2,
	}
	//-------AntGroup type example------
	AntGroup2 = AntGroup{
		NumberOfAnts: 1,
		AntsMap: map[int]Ant{
			1: ant1,
		},
		NumberOfSequence: 2,
	}
	//----------
)

//---------------------------------------------
func TestHello(t *testing.T) {
	//------------
	t.Run("1-an example of Ant struct type", func(t *testing.T) {
		//---Arrange---
		//---Act---
		//---Assert----
		assertIfStruct(t, ant1)
	})
	//------------
	t.Run("2-an example of AntGroup struct type", func(t *testing.T) {
		//---Arrange---
		//---Act---
		//---Assert----
		assertIfStruct(t, AntGroup1)
	})
	// ------------
}

// -----------------------------------------------------
func assertIfStruct(t testing.TB, s interface{}) {
	t.Helper()
	
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		t.Errorf("not struct")
	}

}

// ===============================================
