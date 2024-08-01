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
