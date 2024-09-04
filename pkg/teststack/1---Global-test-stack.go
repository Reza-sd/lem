package teststack

// ====================================
type TestCase struct {
	Skip bool
	Des  string
	Got  any
	Exp  any
}
type TestCasesforFunc struct {
	Skip      bool
	FuncName  string
	TestCases []TestCase
}

//var AllTestCasesPkSlice = []TestCasesFunc{}

// ====================================
