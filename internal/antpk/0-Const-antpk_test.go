package antpk

import (
	"reflect"
	"testing"
)

// ===============================================
/*

 */

// ===============================================
var (
	//-------Ant type example-----------
	ant1 = Ant{
		Name:            "reza",
		CurrentRoomName: "room_1",
		VisitedRoomsArr: []string{"room_0", "room_1"},
		StepNumber:      1,
	}
	//-------Ant type example-----------
	ant2 = Ant{
		Name:            "L3",
		CurrentRoomName: "mio_5",
		VisitedRoomsArr: []string{"mio_7", "mio_3", "mio_5"},
		StepNumber:      2,
	}
	//--------AntGroup type example---------
	antGroup1 = AntGroup{
		NumberOfAnts: 1,
		AntsMap: map[int]Ant{
			1: {Name: "Ant1",
				CurrentRoomName: "room_1",
				VisitedRoomsArr: []string{"room_0", "room_1"},
				StepNumber:      1,
			},
		},
		NumberOfSequence: 2,
	}
	//-------AntGroup type example------
	antGroup2 = AntGroup{
		NumberOfAnts: 1,
		AntsMap: map[int]Ant{
			1: ant1,
		},
		NumberOfSequence: 2,
	}

	//-------AntGroup type example------

	antGroup3 = AntGroup{
		NumberOfAnts: 2,
		AntsMap: map[int]Ant{
			1: ant1,
			2: ant2,
		},
		NumberOfSequence: 2,
	}
	//------------------
	travelPlan1 = TravelPlan{
		FinalSequence: 2,
		TheBestPlan:   antGroup3,
	}
	//------------------
)

// ==================================
func Test_Const_antpk(t *testing.T) {
	//------------
	t.Run("0-an example of non-struct type", func(t *testing.T) {
		//---Arrange---
		expectation := false
		//---Act---
		//---Assert----
		//assertIfStruct(t, "miio") fail test
		assertIfStruct(t, "mio", expectation)
	})
	//------------
	t.Run("1-an example of Ant struct type", func(t *testing.T) {
		//---Arrange---
		expectation := true
		//---Act---
		//---Assert----
		//assertIfStruct(t, "miio") fail test
		assertIfStruct(t, travelPlan1, expectation)
	})
	//------------
	t.Run("2-an example of AntGroup struct type", func(t *testing.T) {
		//---Arrange---
		expectation := true
		//---Act---
		//---Assert----
		assertIfStruct(t, travelPlan1, expectation)
	})
	// ------------

	t.Run("3-an example of TravelPlan struct type", func(t *testing.T) {
		//---Arrange---
		expectation := true
		//---Act---
		//---Assert----

		assertIfStruct(t, travelPlan1, expectation)
	})
	//------------
}

// -----------------------------------------------------
func assertIfStruct(t testing.TB, s interface{}, expectation bool) {
	t.Helper()

	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct && expectation {
		t.Errorf("not struct")
	}

}

// ===============================================
