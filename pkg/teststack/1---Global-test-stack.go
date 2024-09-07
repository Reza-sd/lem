package teststack

// ====================================
type TestCase struct {
	Skip bool
	Des  string
	//Input String //a report of setup generat by fmt.sprintf
	//GotInp struct{Got any,}
	Setup func()(any,string)
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
