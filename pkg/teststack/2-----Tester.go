package teststack

// ====================================

type Tester struct {
	PackageName           string
	AllTestCasesPkSlice   []TestCasesFunc
	AllTestCasesPkStSlice []AllTestCasesPkSt
}

type AllTestCasesPkSt struct {
	Skip          bool
	TestCasesFunc TestCasesFunc
}

// ====================================
