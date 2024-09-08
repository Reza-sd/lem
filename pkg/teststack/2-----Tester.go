package teststack

//====================================
type TestCase struct {
	Skip bool
	Des  string
	Case func() (input string, got any, exp any)
}

type TestCasesforFunc struct {
	Skip      bool
	FuncName  string
	TestCases []TestCase
}

type Tester struct {
	PackageName           string
	AllTestCasesPkStSlice []AllTestCasesPkStruct
}

type AllTestCasesPkStruct struct {
	TestCasesforFunc TestCasesforFunc
	Skip             bool
}
//====================================
