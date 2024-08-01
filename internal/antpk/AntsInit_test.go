package antpk

import (
	
	"testing"
)
//========================================
var (

)
//=========================================
func Test_AntsInit(t *testing.T) {
	//------------
	t.Run("0-an example of non-struct type", func(t *testing.T) {
		//---Arrange---
		// var myantgroup AntGroup

		// ExpectedAntGroup1 := AntGroup{
		// 	NumberOfAnts: 1,
		// 	AntsMap: map[int]Ant{
		// 		1: {Name: "Ant1",
		// 			CurrentRoomName: "room_1",
		// 			VisitedRoomsArr: []string{"room_0", "room_1"},
		// 			StepNumber:      1,
		// 		},
		// 	},
		// 	NumberOfSequence: 2,
		// }
		// numberOfAnts := 1
		//expectation := false
		//---Act---
		//err:= myantgroup.AntsInit(numberOfAnts,myGraph)
		//---Assert----
		//assertIfStruct(t, "miio") fail test
		//assertIfStruct(t, "mio", expectation)
	})
}
//================================================
func assert(t testing.TB, s interface{}, expectation bool) {
	t.Helper()

	// v := reflect.ValueOf(s)
	// if v.Kind() != reflect.Struct && expectation {
	// 	t.Errorf("not struct")
	// }

}
//================================================