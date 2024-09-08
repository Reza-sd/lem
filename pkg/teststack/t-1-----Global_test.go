package teststack

import "testing"

// ========================================
var (
	//-------------
	AllTestCasesPkStSlice2 = []AllCasesPkg{
		//{TestCasesFunc: Init_Test, Skip: true},
		{TestCasesforFunc: Method1_test, Skip: false},
		{TestCasesforFunc: Method2_test},
	}
	//-------------
)

// ========================================
func Test_Tester_pk(t *testing.T) {
	RunAll(t, "pkgname", AllTestCasesPkStSlice2)
}

//========================================
