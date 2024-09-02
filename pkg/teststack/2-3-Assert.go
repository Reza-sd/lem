package teststack

import (
	//"fmt"
	//"fmt"
	"reflect"
	"testing"
)

// =======================================================
func (myTester *Tester) Assert(t testing.TB, got, exp any) {
	t.Helper()

	var gotValue any
	var expValue any

	switch g := got.(type) {
	case func() any:
		gotValue = got.(func() any)() //call function
	default:
		gotValue = g
	}

	switch e := exp.(type) {
	case func() any:
		expValue = exp.(func() any)() //call function
	default:
		expValue = e
	}

	if !reflect.DeepEqual(gotValue, expValue) {
		//t.Fatalf("\n>>>not same: \n got=%v<<\n exp=%v<<", got, exp)
		t.Errorf("\n\n|----> Got=%v<Type=%T\n|----> Exp=%v<Type=%T", gotValue, gotValue, expValue, expValue)
	}

}

//=======================================================
