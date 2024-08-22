package teststack

import (
	"fmt"
	//"reflect"
	"testing"
)

// ===========================
func (myTester *Tester) RunAll(t *testing.T) {
	allTestCasesPkSlice := myTester.AllTestCasesPkStSlice

	for i := 0; i < len(allTestCasesPkSlice); i++ {

		ThisTestCasesFunc := allTestCasesPkSlice[i].TestCasesFunc
		ThisTestCasesFuncSkip := allTestCasesPkSlice[i].Skip

		for j := 0; j < len(ThisTestCasesFunc.TestCases); j++ {
			ThisTestCase := ThisTestCasesFunc.TestCases[j]

			if ThisTestCasesFunc.Skip {
				t.Skip()
			}

			NumfuncDes := fmt.Sprintf("%v/%v-%v", ThisTestCasesFunc.FuncName, j+1, ThisTestCase.Des)

			t.Run(NumfuncDes, func(t *testing.T) {
				if ThisTestCasesFuncSkip {
					t.Skip()
				}

				if ThisTestCase.Skip {
					t.Skip()
				}
				got := ThisTestCase.Got

				exp := ThisTestCase.Exp
				myTester.Assert(t, got, exp)
			})
		}

	}
}

//==============================
