package teststack

// ====================================
// type Case struct {
// 	Skip bool
// 	Des  string
// 	Set  func() (input string, got any, exp any)
// }
// type Case2 struct {
// 	set []func() (string, any, any)
// }

// type AllCasesFunc struct {
// 	//Skip      bool
// 	//FuncName  string
// 	TestCases []Case
// }

type AllCasesPkg struct {
	Name     string
	CasesMap map[string]any
	Skip     bool
}

//====================================
