package teststack

import "testing"

// ========================================
var (
	//-------------
	AllTestCasesPkStSlice2 = []AllCasesPkg{
		//{TestCasesFunc: Init_Test, Skip: true},
		{Name: "func reza", CasesMap: funcReza_Cases, Skip: false},
		{Name: "func reza", CasesMap: funcReza_Cases, Skip: true},
	}
	//-------------
	// //-------------
	// AllTestCasesPkStSlice3 = []AllCasesPkg{
	// 	//{TestCasesFunc: Init_Test, Skip: true},
	// 	{TestCasesforFunc: Method1_test, Skip: false},
	// 	{TestCasesforFunc: Method2_test},
	// }
	// //------------

)

// ========================================
func Test_Tester_pk(t *testing.T) {
	RunAll(t, "pkgname", AllTestCasesPkStSlice2)
	
}

//========================================
