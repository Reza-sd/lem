package teststack

import (
	"fmt"
	//"reflect"
	"testing"
)

// ===========================
func (myTester *Tester) RunOne(t *testing.T, testCaseFunc TestCasesforFunc, skip bool) {

	//--------------------------------------------------
	if skip {
		t.Skip()
	}
	testcaseFuncSlice := testCaseFunc.TestCases
	funcName := testCaseFunc.FuncName

	for j := 0; j < len(testcaseFuncSlice); j++ {
		theTestCase := testcaseFuncSlice[j]
		if theTestCase.Skip {
			t.Skip()
		}

		NumfuncDes := fmt.Sprintf("%v/%v-%v", funcName, j+1, testcaseFuncSlice[j].Des)

		t.Run(NumfuncDes, func(t *testing.T) {

			myTester.Assert(t, theTestCase)
		})
	}

}

//==============================
