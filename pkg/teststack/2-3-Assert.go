package teststack

import (
	//"fmt"
	//"fmt"
	"reflect"
	"testing"
)

// =======================================================
func (myTester *Tester) Assert(t testing.TB, testCase TestCase) {
	t.Helper()

	gotValue,input:=testCase.Setup()

	//got := testCase.Got

	exp := testCase.Exp

	//var gotValue any
	var expValue any

	// switch g := got.(type) {
	// case func() any:
	// 	gotValue = got.(func() any)() //call function
	// default:
	// 	gotValue = g
	// }

	switch e := exp.(type) {
	case func() any:
		expValue = exp.(func() any)() //call function
	default:
		expValue = e
	}

	if !reflect.DeepEqual(gotValue, expValue) {
		//t.Fatalf("\n>>>not same: \n got=%v<<\n exp=%v<<", got, exp)
		t.Errorf("\n\n|---->Input=>>%v<<<\n|---->Got=%v<<<Type of %T>>>\n|---->Exp=%v<<<Type of %T>>>", input,gotValue, gotValue, expValue, expValue)
	}

}

//=======================================================
