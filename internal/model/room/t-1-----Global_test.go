package room

import (

	teststack "main/pkg/teststack"
	"testing"
)

var (
	MyTester = teststack.Tester{
		PackageName:         "mio",
		AllTestCasesPkSlice: []TestCasesFunc{Init_Test},
	}
)

// ========================================
func Test_roompk(t *testing.T) {
	MyTester.Run(t)

}


