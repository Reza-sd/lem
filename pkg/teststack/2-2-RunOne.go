package teststack

import (
	"fmt"
	//"reflect"
	"testing"
)

// ===========================
func (myTester *Tester) RunOne(t *testing.T, testCaseFunc TestCasesFunc) {
	testcaseFuncSlice := testCaseFunc.TestCases
	funcName := testCaseFunc.FuncName
	//for i := 0; i < len(allTestCasesPkSlice); i++ {
	//fmt.Println(mySlice[i])
	for j := 0; j < len(testcaseFuncSlice); j++ {

		if testcaseFuncSlice[j].Skip {
			t.Skip()
		}

		NumfuncDes := fmt.Sprintf("%v/%v-%v", funcName, j+1, testcaseFuncSlice[j].Des)

		t.Run(NumfuncDes, func(t *testing.T) {
			if testcaseFuncSlice[j].Skip {
				t.Skip()
			}
			got := testcaseFuncSlice[j].Got

			exp := testcaseFuncSlice[j].Exp
			myTester.Assert(t, got, exp)
		})
	}

	//}
}

//==============================
