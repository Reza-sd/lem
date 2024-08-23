package teststack

import (
	"testing"
)

var (
	MyTester = Tester{
		//PackageName: "mio",

		AllTestCasesPkStSlice: []AllTestCasesPkSt{
			//{TestCasesFunc: Init_Test, Skip: true},
			{TestCasesFunc: myTester_Cases_1, Skip: false},
			//{TestCasesFunc: myTester_Cases_2},
			//{TestCasesFunc: myTester_Cases_1, Skip: true},
		},
	}
)

// ========================================
func Test_Tester_pk(t *testing.T) {
	MyTester.RunAll(t)
	//MyTester.RunOne(t, Init_Test, true)
	//MyTester.RunOne(t, Init_Test, false)
}
