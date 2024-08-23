package teststack

import (
	//"fmt"
	"fmt"
	"reflect"
	"testing"
)

// =======================================================
func (myTester *Tester) Assert(t testing.TB, got, exp any) (err error) {
	t.Helper()
	//--------------------------------------------------
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("\n\n>>>>>>**** Assert panic ****<<<<<\n%v\n\n", r)
			err = fmt.Errorf("panic: %v", r)

		}
	}()
	//--------------------------------------------------
	var gotValue any
	var expValue any

	switch v := exp.(type) {
	case func() any:
		gotValue = got.(func() any)() //call function
	default:
		gotValue = v
	}

	switch v := exp.(type) {
	case func() any:
		expValue = exp.(func() any)() //call function
	default:
		expValue = v
	}

	if !reflect.DeepEqual(gotValue, expValue) {
		//t.Fatalf("\n>>>not same: \n got=%v<<\n exp=%v<<", got, exp)
		t.Errorf("\n\n|----> got=%v<\n|----> exp=%v<", gotValue, expValue)
	}
	//------------------------------------------------------
	return nil
	//----------------------------------------------------
}

//=======================================================
