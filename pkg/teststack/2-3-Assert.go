package teststack

import (
	"reflect"
	"testing"
)

// =======================================================
func (myTester *Tester) Assert(t testing.TB, testCase TestCase) {
	t.Helper()

	inputs, got, exp := testCase.Case()

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("\n>------------------\nInputs:\n%v\nGot=type:(%T) value:(%v)\nExp=type:(%T) value:(%v)\n>------------------\n",
			inputs, got, got, exp, exp)
	}
}

//=======================================================
