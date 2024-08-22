package teststack

import (
	//"fmt"
	"reflect"
	"testing"
)

// =======================================================
func (myTester *Tester) Assert(t testing.TB, got, exp any) {
	t.Helper()

	var gotValue any
	var expValue any

	switch got.(type) {
	case func() any:
		gotValue = got.(func() any)()
	default:
		gotValue = got
	}

	switch exp.(type) {
	case func() any:
		expValue = exp.(func() any)()
	default:
		expValue = exp
	}

	if !reflect.DeepEqual(gotValue, expValue) {
		//t.Fatalf("\n>>>not same: \n got=%v<<\n exp=%v<<", got, exp)
		t.Errorf("\n>>>not same: \n got=%v<<\n exp=%v<<", got, exp)
	}
}

//=======================================================
