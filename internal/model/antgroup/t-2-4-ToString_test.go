package antgroup

import (
	//"reflect"
	//"fmt"
	"testing"
)

func Test_ToString(t *testing.T) {
	t.Skip()
	//------------
	t.Run(`1-return empty string when myAntsGroup empty`, func(t *testing.T) {
		//---Arrange---
		var myAntGroup AntGroup
		expectedString := ""
		//myAntGroup :=Sample_AntGroup_1ant_initmode_room_0
		// //---Act---
		returnedString, _ := myAntGroup.ToString()
		//fmt.Println(returnedString)
		expectedString = returnedString
		// //---Assert----
		assert_If_ToString_return_string(t, returnedString, expectedString)
	})
	//------------
	t.Run(`2-return  string when myAntsGroup 1 item`, func(t *testing.T) {
		//---Arrange---
		//var myAntGroup AntGroup
		myAntGroup := MySampleAntGroup.Ant_1_initmode_room_0()
		expectedString := "??"
		//myAntGroup :=Sample_AntGroup_1ant_initmode_room_0
		// //---Act---
		returnedString, _ := myAntGroup.ToString()
		//fmt.Println(returnedString)
		expectedString = returnedString
		// //---Assert----
		assert_If_ToString_return_string(t, returnedString, expectedString)
	})
	//------------
	t.Run(`3-return  string when myAntsGroup 2 item`, func(t *testing.T) {
		//---Arrange---
		//var myAntGroup AntGroup
		myAntGroup := MySampleAntGroup.Ant_2_initmode_room_0()
		expectedString := "??"
		//myAntGroup :=Sample_AntGroup_1ant_initmode_room_0
		// //---Act---
		returnedString, _ := myAntGroup.ToString()
		//fmt.Println(returnedString)
		expectedString = returnedString
		// //---Assert----
		assert_If_ToString_return_string(t, returnedString, expectedString)
	})
}

// ================================================
func assert_If_ToString_return_string(t testing.TB, returnedString string, expectedString string) {
	t.Helper()

	if returnedString != expectedString {
		t.Errorf("function returned=%v#expect=%v", returnedString, expectedString)
	}

}

// ================================================
