package teststack

import (
	"fmt"
	"testing"
)

// ===========================
func RunAll(t *testing.T, PackageName string, allTestCasesPkSlice []AllCasesPkg) {
	//--------------------------------------------------
	//allTestCasesPkSlice := myTester.AllTestCasesPkStSlice

	for i := 0; i < len(allTestCasesPkSlice); i++ {

		ThisTestCasesFunc := allTestCasesPkSlice[i].TestCasesforFunc
		ThisTestCasesFuncSkip := ThisTestCasesFunc.Skip
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

				assert(t, &ThisTestCase)

			})
		}

	}
	print("\n")

}

// ==============================
// ====================================
func Inp(inputs ...any) string {
	var str string
	var count int
	for _, item := range inputs {
		count++
		str = str + fmt.Sprintf("%d-inp=type:(%T) value:(%v)\n", count, item, item)
	}
	return str
}

// ====================================
