package teststack

// ====================================
type TestCase struct {
	Skip bool
	Des  string
	got  any
	exp  any
}
type TestCasesFunc struct {
	Skip      bool
	FuncName  string
	TestCases []TestCase
}

var AllTestCasesPkSlice = []TestCasesFunc{}

// ====================================
