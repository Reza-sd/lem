package teststack

import "testing"

// ========================================
var (
	MyTester = Tester{
		PackageName: "mio",

		AllTestCasesPkStSlice: []AllTestCasesPkStruct{
			//{TestCasesFunc: Init_Test, Skip: true},
			{TestCasesforFunc: Method1_test, Skip: false},
			{TestCasesforFunc: Method2_test},
			//{TestCasesFunc: myTester_Cases_1, Skip: true},
		},
	}
)

// ========================================
func Test_Tester_pk(t *testing.T) {
	MyTester.RunAll(t)
}

//========================================
