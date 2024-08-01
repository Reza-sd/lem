package antpk

import (
	"reflect"
	"testing"
)

// ===============================================
/*

 */

// ===============================================
func TestHello(t *testing.T) {
	//------------
	t.Run("1-an example of Ant struct type", func(t *testing.T) {

		//---Arrange---
		ant1 := Ant{
			Name:            "reza",
			CurrentRoomName: "room_1",
			VisitedRoomsArr: []string{"room_0", "room_1"},
			StepNumber:      2,
		}
		//---Act---

		//---Assert----
		assertIfStruct(t, ant1)
	})
	//------------
	t.Run("2-an example of AntGroup struct type", func(t *testing.T) {
		//t.Skip()
		//---Arrange---
		// ant1 := Ant{
		// 	Name:            "reza",
		// 	CurrentRoomName: "room_1",
		// 	VisitedRoomsArr: []string{"room_0", "room_1"},
		// 	StepNumber:      2,
		// }

		AntGroup1 := AntGroup{
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
		//---Act---

		//---Assert----
		assertIfStruct(t, AntGroup1)
	})
	// ------------
}

// -----------------------------------------------------
func assertIfStruct(t testing.TB, s interface{}) {
	t.Helper()
	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Struct {
		t.Errorf("not struct")
	}

}

// ===============================================
