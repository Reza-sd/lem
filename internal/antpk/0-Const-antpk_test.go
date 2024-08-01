package antpk


import (
"testing"
	"reflect"
)

// ===============================================
/*

*/

// ===============================================
func TestHello(t *testing.T) {
	//------------
	t.Run("1-an example of Ant struct type", func(t *testing.T) {

		//---Arrange---
		ant1:=Ant{
			Name: "Reza",
			CurrentRoomName: "room_1",
			VisitedRoomsArr: []string{"room_0","room_1"},
			StepNumber: 2,
		}
		//---Act---

		//---Assert----
		assertIfStruct(t,ant1)
	})
	//------------
	t.Run("2-an example of Ants struct type", func(t *testing.T) {
		//t.Skip()
		//---Arrange---
		ant1:=Ant{
			Name: "Reza",
			CurrentRoomName: "room_1",
			VisitedRoomsArr: []string{"room_0","room_1"},
			StepNumber: 2,
		}
		//---Act---

		//---Assert----
		assertIfStruct(t,ant1)
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