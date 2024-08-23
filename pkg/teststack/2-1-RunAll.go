package teststack

import (
	"fmt"

	"testing"
)

// ===========================
func (myTester *Tester) RunAll(t *testing.T) (err error) {
	//var err error
	//--------------------------------------------------
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("\n\n>>>>>>###### panic RunAll ########<<<<\n%v\n\n", r)
			err = fmt.Errorf("panic: %v", r)

		}
	}()
	//--------------------------------------------------

	allTestCasesPkSlice := myTester.AllTestCasesPkStSlice

	for i := 0; i < len(allTestCasesPkSlice); i++ {

		ThisTestCasesFunc := allTestCasesPkSlice[i].TestCasesFunc
		ThisTestCasesFuncSkip := allTestCasesPkSlice[i].Skip
		print("\n")

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
				errAssert := myTester.Assert(t, got, exp)

				if errAssert != nil {
					t.Errorf("\n\n>>>>>>**** Assert panic ****<<<<<\n%v\n\n", errAssert)
					//return nil
				}
			})
		}

	}
	print("\n")
	return nil
}

//==============================
