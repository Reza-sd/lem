package teststack

// ====================================

type Tester struct {
	//PackageName           string
	AllTestCasesPkStSlice []AllTestCasesPkSt
}

type AllTestCasesPkSt struct {
	TestCasesforFunc TestCasesforFunc
	Skip             bool
}

// ====================================
