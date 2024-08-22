package teststack

// ====================================
type TestCase struct {
	Skip bool
	Des  string
	Got  any
	Exp  any
}
type TestCasesFunc struct {
	Skip      bool
	FuncName  string
	TestCases []TestCase
}

var AllTestCasesPkSlice = []TestCasesFunc{}

// ====================================
