package teststack

import (
	"fmt"
	//"reflect"
	"testing"
)

// ===========================
func (myTester *Tester) RunOne(t *testing.T, testCaseFunc TestCasesforFunc, skip bool) {
	//--------------------------------------------------
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("\n\n>>>>>>###### panic RunOne ########<<<<\n%v\n\n", r)

			//err = fmt.Errorf("panic: %v", r)

		}
	}()
	//--------------------------------------------------
	if skip {
		t.Skip()
	}
	testcaseFuncSlice := testCaseFunc.TestCases
	funcName := testCaseFunc.FuncName
	//for i := 0; i < len(allTestCasesPkSlice); i++ {
	//fmt.Println(mySlice[i])
	for j := 0; j < len(testcaseFuncSlice); j++ {
		theTestCase := testcaseFuncSlice[j]
		if theTestCase.Skip {
			t.Skip()
		}

		NumfuncDes := fmt.Sprintf("%v/%v-%v", funcName, j+1, testcaseFuncSlice[j].Des)

		t.Run(NumfuncDes, func(t *testing.T) {
			// if testcaseFuncSlice[j].Skip {
			// 	t.Skip()
			// }
			// _,input:=theTestCase.Setup()

			// got := theTestCase.Got

			// exp := theTestCase.Exp
			myTester.Assert(t, theTestCase)
		})
	}

	//}
	//return nil
}

//==============================
