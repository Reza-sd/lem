package teststack

// ====================================

type Tester struct {
	//PackageName           string
	AllTestCasesPkStSlice []AllTestCasesPkSt
}

type AllTestCasesPkSt struct {
	TestCasesFunc TestCasesFunc
	Skip          bool
}

// ====================================
