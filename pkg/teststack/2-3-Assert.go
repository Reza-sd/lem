package teststack

import (
	"reflect"
	"testing"
)

// =======================================================
func assert(t testing.TB, testCase *Case) {
	t.Helper()

	inputs, got, exp := testCase.Set()

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("\n>------------------\nInputs:\n%v\nGot=type:(%T) value:(%v)\nExp=type:(%T) value:(%v)\n>------------------\n",
			inputs, got, got, exp, exp)
	}
}

//=======================================================
