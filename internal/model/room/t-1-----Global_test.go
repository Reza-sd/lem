package room

import (
	teststack "main/pkg/teststack"
	"testing"
)

var (
	MyTester = teststack.Tester{
		PackageName: "mio",
		//AllTestCasesPkSlice: []TestCasesFunc{Init_Test},
		AllTestCasesPkStSlice: []AllTestCasesPkSt{{
			Skip:          false,
			TestCasesFunc: Init_Test,
		}},
	}
)

// ========================================
func Test_roompk(t *testing.T) {
	MyTester.RunAll(t)
	//MyTester.RunOne(t, Init_Test, true)
	//MyTester.RunOne(t, Init_Test, false)
}

// func Test_Init(t *testing.T) {
// 	//MyTester.RunAll(t)
// 	//MyTester.RunOne(t, Init_Test, true)
// 	MyTester.RunOne(t, Init_Test, false)
// }
// ========================================
