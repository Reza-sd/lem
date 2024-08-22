package teststack

import (
	"fmt"
	//"reflect"
	"testing"
)

// ===========================
func (myTester *Tester) RunAll(t *testing.T) {
	allTestCasesPkSlice := myTester.AllTestCasesPkSlice

	for i := 0; i < len(allTestCasesPkSlice); i++ {
		//fmt.Println(mySlice[i])
		for j := 0; j < len(allTestCasesPkSlice[i].TestCases); j++ {

			if allTestCasesPkSlice[i].Skip {
				t.Skip()
			}

			NumfuncDes := fmt.Sprintf("%v/%v-%v", allTestCasesPkSlice[i].FuncName, j+1, allTestCasesPkSlice[i].TestCases[j].Des)

			t.Run(NumfuncDes, func(t *testing.T) {
				if allTestCasesPkSlice[i].TestCases[j].Skip {
					t.Skip()
				}
				got := allTestCasesPkSlice[i].TestCases[j].Got

				exp := allTestCasesPkSlice[i].TestCases[j].Exp
				myTester.Assert(t, got, exp)
			})
		}

	}
}

//==============================
