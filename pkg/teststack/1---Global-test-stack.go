package teststack

import (
	"fmt"
)

// ====================================
type TestCase struct {
	Skip bool
	Des  string
	//Input String //a report of setup generat by fmt.sprintf
	//GotInp struct{Got any,}
	Case func() (input string, got any, exp any)
	//Got  any
	//Exp  any
}
type TestCasesforFunc struct {
	Skip      bool
	FuncName  string
	TestCases []TestCase
}

//var AllTestCasesPkSlice = []TestCasesFunc{}

// ====================================
func Inp(inputs ...any) string {
	var str string
	var count int
	for _, item := range inputs {
		count++
		str = str + fmt.Sprintf("%d-inp=type:(%T) value:(%v)\n", count, item, item)
	}
	return str
}

// ====================================
